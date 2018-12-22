from __future__ import print_function
import os
import shutil
import subprocess
# Get the root project directory
PROJECT_DIRECTORY = os.path.realpath(os.path.curdir)

def remove_file(filename):
    fullpath = os.path.join(PROJECT_DIRECTORY, filename)
    if os.path.exists(fullpath):
        os.remove(fullpath)

def remove_dir(dirname):
    shutil.rmtree(os.path.join(
        PROJECT_DIRECTORY, dirname
    ))

def remove_all_files(filename):
    for name in [filename,]:
        os.remove(os.path.join(
            PROJECT_DIRECTORY, name
        ))

if '{{ cookiecutter.use_docker }}'.lower() != 'y':
    remove_all_files("Dockerfile")

if '{{ cookiecutter.use_tls }}'.lower() != 'y':
    remove_dir("certs/")

if '{{ cookiecutter.use_ci}}'.lower() != 'circle':
    remove_dir(".circleci/")

    init = subprocess.Popen("make init", shell=True, stderr=subprocess.STDOUT)
