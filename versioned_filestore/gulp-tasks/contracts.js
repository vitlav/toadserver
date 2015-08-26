var gulp = require('gulp');
var gDebug = require('gulp-debug');
var gUtil = require('gulp-util');
var path = require('path');
var del = require('del');
var fs = require('fs-extra');
var gsm = require('gulp-smake');
var os = require('os');

var exports = {};

var options = require('../contracts/contracts.json');
var build = options.build;
var buildtests = options.buildtests;

var buildDir = path.join(build.root, build.buildDir);
var docsDir = path.join(build.root, build.docsDir);

// Just some debugging info. Enable if the SOL_UNIT_BUILD_DEBUG envar is set.
var debugMode = true;// process.env.SOL_UNIT_BUILD_DEBUG;
var dbg;

if(debugMode){
    dbg = gDebug;
} else {
    dbg = gUtil.noop;
}

// TODO start separating tasks into different files.

// Removes the build folder.
gulp.task('contracts-clean', function(cb) {
    del([buildDir, docsDir], cb);
});

// Used to set up the project for building.
gulp.task('contracts-init-build', function (cb) {
    del([buildDir, docsDir], cb);
});

// Writes the complete source tree to a temp folder. This is also where external source dependencies
// would be fetched and set up if needed.
gulp.task('contracts-pre-build', ['contracts-init-build'], function(){
    // Create an empty folder in temp to use as temporary root when building.
    var temp = path.join(os.tmpdir(), "sol-unit");
    fs.emptyDirSync(temp);
    // Modify the source folder in the object, so that it uses the new temp folder.
    exports.base = temp;
    // Create the path to the root source folder.
    var base = path.join(build.root, build.sourceDir);
    return gulp.src(build.paths, {base: base})
        .pipe(dbg())
        .pipe(gulp.dest(temp));
});

// Compiles the contracts. This is also where external code dependencies would be set up and built if needed.
gulp.task('contracts-build', ['contracts-pre-build'], function () {
    return gulp.src(exports.base + '/**/*')
        .pipe(dbg())
        .pipe(gsm.build(build, exports));
});

// Writes the complete source tree to a temp folder. This is also where external source dependencies
// would be fetched and set up if needed.
gulp.task('contracts-pre-build-tests', ['contracts-pre-build'], function(){

    // Create the path to the root source folder.
    var base = path.join(build.root, buildtests.testDir);
    return gulp.src(buildtests.paths, {base: base})
        .pipe(dbg())
        .pipe(gulp.dest(exports.base));
});

// Compiles the contracts. This is also where external code dependencies would be set up and built if needed.
gulp.task('contracts-build-tests', ['contracts-pre-build-tests'], function () {
    exports.buildDir = path.join(build.root, buildtests.buildDir);
    return gulp.src(exports.base + '/**/*')
        .pipe(dbg())
        .pipe(gsm.build(build, exports));
});

// Any cleanup of the build directory is put here.
gulp.task('contracts-post-build-tests', ['contracts-build-tests'], function(cb){
    del([path.join(buildDir,'Asserter.*'), path.join(docsDir,'Asserter.*')], cb);
});