# go-boid-go

Note: This is not a beginner friendly repository. 

So what this seemingly green dots on a random terminal screen, all about? <br>
First of all, those random green dots are **BOIDs (Android Birds)**. 

This project is an attempt to learn multi-threading in golang. 
A thread is a tool or abstraction that allows us to perform parallel computation. (remember your OS lectures?)

There are some theorems we will be using to this program, indirectly. 
1. Amdahl's Law (you cannot keep on speeding up a work/process (only 1) to infinity by simply enough workers/processors at it, it reaches to a finite limit.)
2. Gustafson's Law (same as Amdahl's law but here you have more than 1 work/processes to do at a time)


Some terminology:-
1. Processes - You start 2 instances(windows) of notepad.exe, they both are separate processes. They are not sharing resources. One does not affect the other. Isolated. They are a bit heavy on resources relatively. E.g - fork()

Go focuses more on threads and green threads. 

2. Threads(Kernel Level Thread) - A thread is the solution to downsides of a process. Threads share memory b/w them, unlike processes. Faster to create. Threads are not isolated. 
3. Green Thread(User Level Threads) - Efficient version of thread. It tries to reduce the time in context switching of threads when an interruption happens and new process/thread from ready queue enters into processing stage. OS picks a new thread from ready queue when the interruption occurs. Reduces context switch overhead. Helpful when there are a lot of processes on ready queue. 


User level thread runs inside kernel level thread. But a green thread has some disadvantages. Hence, Go uses a mixture of Green threads running under normal threads. Go reshuffles green threads as and when required, according to priority of task.

Hence, our BOIDs are actually threads, implemented by goroutines. 

One boid = One thread. 







Final preview of what it looks like at the end. 

Why this?

What is this?

What it teaches?

Answer to all these question will be updated soon!

![](https://user-images.githubusercontent.com/52788043/151409741-ed72f516-ac12-43f5-905e-bd6aa39c2194.gif)

![](https://user-images.githubusercontent.com/52788043/151409751-9bdb3163-6160-4f96-a43b-34c927eb3dde.gif)

This 2D simulation tried to mimic flocking behaviour of large group of birds flying. 

Concept of Boids are in the introducture lecture of A-Life(Artificial Life) classes

Graphics Library used - EBITEN. 

Each Boid has an velocity vector, position vector and Id.


