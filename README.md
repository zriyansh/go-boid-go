# go-boid-go

Note: This is not a beginner friendly Go repository. 

Final preview of what it looks like at the end. 

![](https://user-images.githubusercontent.com/52788043/151409751-9bdb3163-6160-4f96-a43b-34c927eb3dde.gif)

This 2D simulation tried to mimic flocking behaviour of large group of birds flying. 

Concept of Boids are in the introducture lecture of A-Life(Artificial Life) classes

Graphics Library used - EBITEN. 

<hr>

So what this seemingly green dots on a random terminal screen, all about? <br>
First of all, those random green dots are **BOIDs (Android Birds)**. 

This project is an attempt to learn multi-threading in golang. 
A thread is a tool or abstraction that allows us to perform parallel computation. (remember your OS lectures?)

<hr>

There are some theorems we will be using to this program, indirectly. 
1. Amdahl's Law (you cannot keep on speeding up a work/process (only 1) to infinity by simply enough workers/processors at it, it reaches to a finite limit.)
<img width="418" src="https://user-images.githubusercontent.com/52788043/151524834-75e6dade-df47-46e4-9dc5-5c20c100d051.png">

2. Gustafson's Law (same as Amdahl's law but here you have more than 1 work/processes to do at a time)
<img width="479" src="https://user-images.githubusercontent.com/52788043/151524847-e981eb45-e024-4b6f-94d6-433954dda503.png">


<hr>

Some terminology:-
1. Processes - You start 2 instances(windows) of notepad.exe, they both are separate processes. They are not sharing resources. One does not affect the other. Isolated. They are a bit heavy on resources relatively. E.g - fork()

Go focuses more on threads and green threads. 

2. Threads(Kernel Level Thread) - A thread is the solution to downsides of a process. Threads share memory b/w them, unlike processes. Faster to create. Threads are not isolated. 
3. Green Thread(User Level Threads) - Efficient version of thread. It tries to reduce the time in context switching of threads when an interruption happens and new process/thread from ready queue enters into processing stage. OS picks a new thread from ready queue when the interruption occurs. Reduces context switch overhead. Helpful when there are a lot of processes on ready queue. 

<hr>

User level thread runs inside kernel level thread. But a green thread has some disadvantages. Hence, Go uses a mixture of Green threads running under normal threads. Go reshuffles green threads as and when required, according to priority of task.

Hence, our BOIDs are actually threads, implemented by goroutines. 

One boid = One thread. 

Each Boid has an velocity vector, position vector and Id.

There are 2 types of Inter-Process Communication(IPC) that can happen between processes. 
1. Message Passing
2. Shared Memory

We define a ```viewRadius``` of a Boid (eyes of the boid) to observe which all other boids are in sight. 

<img width="694" src="https://user-images.githubusercontent.com/52788043/151524639-f7f2a6e8-0c3a-4149-a0ed-0adc7e0530ed.png">

<hr>

3 properties of boids that we will define using memory sharing. 
1. Alignment - Align boids which come under the viewRadius
2. Cohesion - keep them in a group
3. Separation - do not let them merge together into a single boid. 

Every boid is aware of position and velocity of other boids, using 2D array initialized to -1 (means no boid in that 2D frame), when a boid is found, we replace -1 with boidId. 

Boids in the viewRadius will adapt to one another's velocity and move in a group. (Alignment)

Now we need a proper thread syncronozation (using mutexes) so that race condition do not occur. That is, boids shold not update their position when we scan them, else they might be counted twice, or not at all. 


```
// defined as a variable
lock = sync.Mutex{}

// lock the process when it starts
lock.Lock() 

// code //

// unlock it at end to be used by other processes
lock.Unlock()
```

But this normal lock and unlock system is not efficient, so we use readers and writers lock system. 

Then we introduce a bounce mechanism that mimics real life birds when they see a wall and slows down and changes the direction.

<img width="604" src="https://user-images.githubusercontent.com/52788043/151524569-4905cec5-1646-4c01-8ac8-42f3285619ec.png">


# How to run this simulation in your system 

This is fairly simple codebase with 














