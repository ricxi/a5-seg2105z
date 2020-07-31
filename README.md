# Assignment 5 (SEG 2105 Z)

Pandemic Registration System (A component based on Assignment #2)  

**Description:** a simple system that patients can use to register for a virus test. Medical staff can then open a repository  

Code for this system is stored under the directory pantracker: <https://github.com/irixoc/a5-seg2105z/tree/master/pantracker>  

## Design & Architecture

### Design Decisions
The system was constrcuted using a mixture of **top-down** and **bottom-up** design, but with a more central focus on bottom-up design.
I knew I wanted the overall architecture to be MVC, so I built the small individual components first. Afterwards, I would start arranging, organizing, and adjusting sections of the system to fit into the MVC architecture.

### Tech Stack
A list of all the programming languages, technology, and frameworks that I used:  
* Go 1.14
    * Gorilla Web Framework for Go
* CSS
    * Bootstrap (CSS Framework)
    * JQuery Framework required as a depedency for Bootstrap
* HTML

### Design Principles
Various design principles (Chapter 9A) learned in class were used. 

1. **Divide and Conquer**  
Invidiual components were built separately and then put together. For example, the patient model was built, then the patient controller, and patient view were built. Finally, all the layers were connected together. 

2. **Increase Cohesion**  
* I tried to use as much **Functional Cohesion** as possible. I was able to achieve this breaking up complicated logic into more methods and functions that returned one value or performed fewer side-effects. However, this was not always possible, as I had to update the make-shift database I had (the json file). 
* **Utility Cohesion** was used to store certain functions in methods that were used by many different subsystems, such as a function that read in JSON information.

3. **Decoupling Design:**
* Understanding the trade-off between **Stamp Coupling** and **Data Coupling**, I chose to *reduce* Data Coupling. This choice was based on the fact that Go did not have many complicated data types, and it was also the standard in Go to pass less arguments over having slightly more complicated arguments (e.g., passing a struct).  
* **Common Coupling** was completely eliminated by storing global variables in structs, which is a practice that is also common when it comes to web development in Go. There is also a technique to implement the **Singleton** design pattern in Go, which may be integrated in later iterations.  

### 1st Iteration: Building a Minimum Viable System (based on the following user stories)
#### User Stories
1. User enters personal information on patient information page
2. Test Provider can view available users on a patient repository and select a user for a test
3. Lab Technician accesses patient repository to record test results
4. User can login and see their test results

State will be stored using a JSON file instead of an SQL database. This was done to simplify things so that the application can be built quickly with less overhead. For example, all patients and their information will be stored in a json file. The plan is to switch to SQL for later iterations. The system was built in accordance to the principles of **Stable Architecture**, so the entire models package can be **reeingineered** to use an sql database without affecting the system.

### 2nd Iteration: Pivoting for time (changing user stories)  
As with most software projects, sometimes there isn't enough time to finish it. Thus, I decided to pivot and build a simpler system. I updated the user stories to reflect this new system.
1. Patient enters personal information to register for a viral test.
2. Medical staff opens a repository to view all registered patients and book an appointment for them.


### A note about the Go struct
The structs in Go look a lot like structs in C. However, they are a bit more versatile, but are used in a way that may look confusing to some. I am including a short explanation here of how they are used because I also found it a bit confusing in the beginning when I was researching on how to build a web app in Go.
1. Methods can be added to structs in go, so OOP design can be used.  
For example, PatientDB is a struct, but it has a constructor called NewPatientDB, and it also has a method called CreatePatient. Methods for a struct have this syntax: *func(s *StructName) MethodName (arguments)*  
![OOPstyle](https://github.com/irixoc/a5-seg2105z/blob/master/rmImages/OOP2.png)

2. A struct can also be used to work easily with JSON objects  
For those types of structs, they'll often have a label beside them (as seen below)  
![jsonStyle](https://github.com/irixoc/a5-seg2105z/blob/master/rmImages/json.png)  

### Additional Notes
Public methods in go start with a captial letter.  
Private methods start with a lower case letter.
