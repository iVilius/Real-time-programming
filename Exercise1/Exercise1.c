// gcc 4.7.2 +
// gcc -std=gnu99 -Wall -g -o Exercise1 Exercise1.c -lpthread

#include <pthread.h>
#include <stdio.h>
#include <sys/types.h>

int i = 0;
pthread_mutex_t lock;

void* thread_1(){
	
	int j;
	for (j = 0; j < 1000000; j++){
		pthread_mutex_lock(&lock);
		i++;
		pthread_mutex_unlock(&lock);
	}
	
	return NULL;
}


void* thread_2(){
	
	int j;
	for (j = 0;j < 1000001; j++){
		pthread_mutex_lock(&lock);
		i--;
		pthread_mutex_unlock(&lock);
	}
	
	return NULL;
}

int main(){
	pthread_mutex_init(&lock, NULL);
	pthread_t someThread1;
	pthread_t someThread2;
	pthread_create(&someThread1, NULL, thread_1, NULL);
	pthread_create(&someThread2, NULL, thread_2, NULL);
	
	pthread_join(someThread1, NULL);
	pthread_join(someThread2, NULL);
	pthread_mutex_destroy(&lock);	

	printf("%d \n", i);

	return 0;
}
