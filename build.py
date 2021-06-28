import os, platform, sys, shutil, glob


main_dir = os.path.dirname(os.getcwd())
backend_path = "AstralTest"

if len(sys.argv) < 2:
    print("outputDir is not found")
    sys.exit(0)

output_dir = os.path.join(main_dir, sys.argv[1])

def set_linux_env():
    os.environ['GOOS'] = "linux"
    os.environ['GOARCH'] = "amd64"

def set_win_env():
    os.environ['GOOS'] = "windows"
    os.environ['GOARCH'] = "amd64"

def build_package(input_dir, system):
    print("Start building the project")
    if system == "linux":
        set_linux_env()
    elif system == "windows":
        set_win_env()
    elif system == "raspberry":
        set_raspberry_env()
    #cd to main directory
    os.chdir(main_dir)
    #if build directory doesn't exist then create one
    try:
        os.stat(output_dir)
    except:
        os.mkdir(output_dir) 
    #recieve package name from input path
    if platform.system() == "Windows":
        lst = input_dir.split("\\")
    else:    
        lst = input_dir.split("/")
    package_name = lst[len(lst)-1]
    p = os.path.join(output_dir, package_name)   
    #if package folder in building folder doesn't exist then create one
    try:
        os.stat(p)
    except:
        os.mkdir(p)
    #cd to package directory
    os.chdir(input_dir)
    #run go build to created directory
    dest = os.path.join(output_dir, package_name)
    os.system(" GO111MODULE=off go build -o " + dest)
    os.chdir(main_dir)
    print("Package is built")    

def copy_docker_files(input_dir, file_names:list):
    print("Copy docker files")
    os.chdir(input_dir)
    for f in file_names:
       shutil.copyfile(os.path.join(input_dir, f), os.path.join(output_dir, f))
    os.chdir(main_dir)      

def build_all(os_to_build, docker_files:list):
    build_package(os.path.join(main_dir, backend_path), os_to_build)
    copy_docker_files(os.path.join(main_dir, backend_path, "docker"), docker_files)


docker_files = ["docker-compose.yml", "init_db.sql", "nginx.conf"]
build_all(sys.argv[2], docker_files)
