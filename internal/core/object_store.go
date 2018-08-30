package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ObjectStore interface {
	Get(string, ...string) ([]byte, error)
	Put([]byte, string, ...string) error
}

type FileSystemStore struct {
	basedir string
}

func NewFilesystemStore(basedir string) *FileSystemStore {
	absdir, _ := filepath.Abs(basedir)
	return &FileSystemStore{
		basedir: absdir,
	}
}

func (fs *FileSystemStore) Get(path string, subpath ...string) ([]byte, error) {
	if filepath.IsAbs(path) {
		return nil, fmt.Errorf("get object: only relative paths allowed")
	}

	// go doesn't allow passing mixed values and variadic arrays so we need to
	// collapse it to a single array first
	keys := append([]string{fs.basedir, path}, subpath...)
	objectPath := filepath.Join(keys...)

	return ioutil.ReadFile(objectPath)
}

func (fs *FileSystemStore) Put(data []byte, path string, subpath ...string) error {
	// FIXME: This should transparently create directories when required, to be
	// consistent with the behaviour of s3
	if filepath.IsAbs(path) {
		return fmt.Errorf("get object: only relative paths allowed")
	}

	// go doesn't allow passing mixed values and variadic arrays so we need to
	// collapse it to a single array first
	keys := append([]string{fs.basedir, path}, subpath...)
	objectPath := filepath.Join(keys...)

	return ioutil.WriteFile(objectPath, data, 0644)
}

type S3ObjectStore struct {
	client *s3.S3
	bucket string
	prefix string
}

func NewS3ObjectStore(bucket, prefix string) *S3ObjectStore {
	sess := session.Must(session.NewSession())
	client := s3.New(sess)
	return &S3ObjectStore{
		client: client,
		bucket: bucket,
		prefix: prefix,
	}
}

func (store *S3ObjectStore) Get(objpath string, subpath ...string) ([]byte, error) {
	// go doesn't allow passing mixed values and variadic arrays so we need to
	// collapse it to a single array first
	keys := append([]string{store.prefix, objpath}, subpath...)
	key := path.Join(keys...)
	obj, err := store.client.GetObject(
		&s3.GetObjectInput{
			Bucket: &store.bucket,
			Key: &key,
		},
	)
	if err != nil {
		return nil, err
	}
	defer obj.Body.Close()

	return ioutil.ReadAll(obj.Body)
}

func (store *S3ObjectStore) Put(data []byte, objpath string, subpath ...string) error {
	buf := bytes.NewReader(data)
	// go doesn't allow passing mixed values and variadic arrays so we need to
	// collapse it to a single array first
	keys := append([]string{store.prefix, objpath}, subpath...)
	key := path.Join(keys...)
	_, err := store.client.PutObject(
		&s3.PutObjectInput{
			Bucket: &store.bucket,
			Key: &key,
			Body: buf,
		},
	)
	if err != nil {
		return fmt.Errorf("put object: %v", err)
	}

	return nil
}

