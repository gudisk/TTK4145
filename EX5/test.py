from ctypes import cdll

test_lib = cdll.LoadLibrary("test.so")
test = test_lib.test

def main():
	print test()

main()