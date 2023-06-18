# you can add this file into your Golang build configuration,
# so every time you build the project, it will start the
# redis automatically.
import os
import platform

os_name = platform.system()
if os_name == "Windows":
    if os.system("redis-server.exe") != 0:
        print("fail to run redis")
elif os_name == "Linux" or os_name == "Darwin":
    print("I haven't write the script for ", os_name)
