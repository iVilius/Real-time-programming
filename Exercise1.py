from threading import Thread
import threading # why cannot import threading from threading??

i = 0
Lock = threading.Lock()

def thread_1():
	
	global i
	for j in range(0,1000000):
		Lock.acquire()
		i = i + 1
		Lock.release()

def thread_2():

	global i
	for k in range(0,1000001):
		Lock.acquire()
		i = i - 1
		Lock.release()
	


def main():
	someThread1 = Thread(target = thread_1, args = (),)
	someThread2 = Thread(target = thread_2, args = (),)
	
	someThread1.start()
	someThread2.start()

	someThread1.join() 
	someThread2.join()

	print i

main()
