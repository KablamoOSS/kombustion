var gulp = require('gulp');
var sass = require('gulp-sass');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var rename = require('gulp-rename');
var minify = require('gulp-clean-css');

var scripts = [
  "static-src/_js/Core/Kube.js",
  "static-src/_js/Core/Kube.Plugin.js",
  "static-src/_js/Core/Kube.Animation.js",
  "static-src/_js/Core/Kube.Detect.js",
  "static-src/_js/Core/Kube.FormData.js",
  "static-src/_js/Core/Kube.Response.js",
  "static-src/_js/Core/Kube.Utils.js",
  "static-src/_js/Message/Kube.Message.js",
  "static-src/_js/Sticky/Kube.Sticky.js",
  "static-src/_js/Toggleme/Kube.Toggleme.js",
  "static-src/_js/Offcanvas/Kube.Offcanvas.js",
  "static-src/_js/Collapse/Kube.Collapse.js",
  "static-src/_js/Dropdown/Kube.Dropdown.js",
  "static-src/_js/Tabs/Kube.Tabs.js",
  "static-src/_js/Modal/Kube.Modal.js",
  "static-src/_js/custom.js"
];

gulp.task('sass', function() {
    return gulp.src('static-src/kube.scss')
        .pipe(sass())
        .pipe(gulp.dest('static/css'))
        .pipe(rename('kube.min.css'))
        .pipe(minify())
        .pipe(gulp.dest('static/css'));
});

gulp.task('combine', function() {
    return gulp.src([
            'static-src/_scss/_variables.scss',
            'static-src/_scss/mixins/_breakpoints.scss',
            'static-src/_scss/mixins/_fonts.scss',
            'static-src/_scss/mixins/_flex.scss',
            'static-src/_scss/mixins/_grid.scss',
            'static-src/_scss/mixins/_utils.scss',
            'static-src/_scss/mixins/_buttons.scss',
            'static-src/_scss/mixins/_gradients.scss',
            'static-src/_scss/mixins/_labels.scss'
        ])
        .pipe(concat('kube-compiled.scss'))
        .pipe(gulp.dest('static-src'));
});

gulp.task('scripts', function() {
    return gulp.src(scripts)
        .pipe(concat('kube.js'))
        .pipe(gulp.dest('static/js'))
        .pipe(rename('kube.min.js'))
        .pipe(uglify())
        .pipe(gulp.dest('static/js'));
});

gulp.task('watch', function() {
    gulp.watch(scripts, ['scripts']);
    gulp.watch([
      'static-src/_scss/*.scss',
      'static-src/_scss/components/*.scss',
      'static-src/_scss/mixins/*.scss',
      'static-src/_scss/custom/*.scss'
    ],
    [
      'combine',
      'sass'
    ]);

});

gulp.task('default', ['sass', 'combine', 'scripts',  'watch']);
