import sys
import time

if __name__ == "__main__":
    print("Hello, World!")
    print(f"Arguments: {sys.argv[1:]}")
    time.sleep(5)
    print("Hello, World! 5s later")