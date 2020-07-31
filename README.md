# Assignment 5 (SEG 2105 Z)

Extension/Revamp of the Pandemic System (Assignment #2)  

**Description:** a simple system that helps test providers to administer viral tests and for lab technicians to record test results for patients.  


## Design & Architecture

### Design Decisions
The system was constrcuted using a mixture of **top-down** and **bottom-up** design, but with a more central focus on bottom-up design.
I knew I wanted the overall architecture to be MVC, so I built the small individual components first, and I would start arranging, organizing, and adjusting sections of the system to be fit into the MVC architecture.

### Design Principles
Various design principles (Chapter 9A) learned in class were used. 

1. **Divide and Conquer**  
Invidiual components were built separately and then put together. For example, the patient model was built, then the patient controller, and patient view were built. Finally, all the layers were connected together. 

2. **Increase Cohesion**
**Utility Cohesion** was used to store certain functions in methods that were used by many different subsystems, such as a function that read in JSON information.

3. **Decoupling Design:** 
Understanding the trade-off between **Stamp Coupling** and **Data Coupling**, I chose to *reduce* Data Coupling. This choice was based on the fact that Go did not have many complicated data types, and it was also the standard in Go to pass less arguments over having slightly more complicated arguments (e.g., passing a struct).  

**Common Coupling** was completely eliminated by storing global variables in structs, which is a practice that is also common when it comes to web development in Go. There is also a technique to implement the **Singleton** design pattern in Go, which may be integrated in later iterations.  

### 1st Iteration: build a minimum viable system based on the following user stories
#### User Stories
1. User enters personal information on patient information page
2. Test Provider can view available users on a patient repository and select a user for a test
3. Lab Technician accesses patient repository to record test results
4. User can login and see their test results

State will be stored using a JSON file instead of an SQL database. For example, all patients and their information
will be stored in a json file. The plan is to switch to SQL in later iterations. The entire models package can be reeingineered to use
an sql database without affecting the system


### Tech Stack
A list of all the programming languages, technology, and frameworks that I used:  
* Go 1.14 (no front/back-end web framework because go comes with built-in a web server and web templating library)
* CSS
* HTML
* Bootstrap (CSS Framework)

### A note about the Go struct
The structs in Go look a lot like structs in C. However, they are a bit more versatile, but may be used in a way that may look confusing. I am including a short explanation here of how they are used because it also tripped me out in the beginning when I was researching on how to build a web app in Go.
1. Methods can be added to structs in go, so OOP design can be used
![OOPstyle](https://github.com/irixoc/a5-seg2105z/blob/master/rmImages/OOP2.png)
2. structs can also be used to work with JSON data
![jsonStyle](https://github.com/irixoc/a5-seg2105z/blob/master/rmImages/json.png)  

In the project, there are going to be a few different usages of structs.  
structs used to create and store 