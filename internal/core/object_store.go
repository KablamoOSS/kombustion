package core

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
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
	var keys []string

	// go doesn't allow passing mixed values and variadic arrays so we need to
	// collapse it to a single array first
	if filepath.IsAbs(path) {
		keys = append([]string{path}, subpath...)
	} else {
		keys = append([]string{fs.basedir, path}, subpath...)
	}

	objectPath := filepath.Join(keys...)

	data, err := ioutil.ReadFile(objectPath)
	if os.IsNotExist(err) {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("get object: %v", err)
	}
	return data, nil
}

func (fs *FileSystemStore) Put(data []byte, path string, subpath ...string) error {
	var keys []string

	// go doesn't allow passing mixed values and variadic arrays so we need to
	// collapse it to a single array first
	if filepath.IsAbs(path) {
		keys = append([]string{path}, subpath...)
	} else {
		keys = append([]string{fs.basedir, path}, subpath...)
	}

	objectPath := filepath.Join(keys...)

	// Transparently create directories when required, to be
	// consistent with the behaviour of s3
	dirName := filepath.Dir(objectPath)
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		return fmt.Errorf("put object: create path: %v", err)
	}

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
			Key:    &key,
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
			Key:    &key,
			Body:   buf,
		},
	)
	if err != nil {
		return fmt.Errorf("put object: %v", err)
	}

	return nil
}
