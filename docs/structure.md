# IOT HVA

# Course Info

![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled.png)

## Website

[](https://dlo.mijnhva.nl/d2l/le/lessons/471080/lessons/1556473)

- Bewertet basierend auf Verbesserung und Erwartungen
- Konsistens ist key
- Basiert auf verschiedenen Kategorien: web, itsec, embedded, business, ui
- Alles muss im Portfolio mit Markdown sein
- Update portfolio weekly
- mysql workbench

## Blueprint structure and requirements

### Infos

[weather-station.html](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/weather-station.html)

### Project organisation

- Connection between Notion and Markdown Gitlab server?!? → https://github.com/souvikinator/notion-to-md bash script which uploads daily an updated version of my notion as markdown to the documentation (bash script or command alias)
- Issue managment?
- Can be developed in Github as fork and later merge it
- Should everything be done at the same time? because weekly commits are wanted or is that because we use the gitlab documentation
- Are there example repo’s?

### Timeline

- ************************************************************************Basic should be finished in 3 weeks************************************************************************
- First get arduino with minimum sensors running
- Build Backend to fetch from client
- Build and design 3D chasis (Because website can be done at home)
- Create website to display data from backend

### Backend

- Docker container with NGINX and MySQL
- User authentication
- Endpoints
- Security:
    - Https
    - SSL?
    - User login
    - Encryption → Check with ChatGPT for Algos and Architecture
- Punktevergabe Backend
    
    ![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%201.png)
    
- Punktevergabe Database
    
    ![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%202.png)
    

### Frontend

- No inline stuff
- Tailwind css??
- Displaying values with → **Chart.js**
- AJAX
- Punktevergabe Frontend
    
    ![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%203.png)
    

### Embedded

- Structure in Fritzing
- Input Sensonrs
    - Temperature
    - Wind?
    - Humidity
- Output
    - Display for basic information → details on website
    - Speaker for battery problems
- Solar and Wind powered
- LoRa?
- Where are the criterias ?!?!?

### Physical Design

- 3d printed modular casing to allow more expansions later
    - each sensor one module (or similar sensors one module)
    - standardized mounting points
    - central mounting bar (tree structure)
- Or pipe system

![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%204.png)

![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%205.png)

- Punktevergabe Design
    
    ![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%206.png)
    
    ![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%207.png)
    
    ![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%208.png)
    
- Punktevergabe Create
    
    ![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%209.png)
    
    ![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%2010.png)
    
    ![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%2011.png)
    
- Punktevergabe Failures
    
    ![Untitled](IOT%20HVA%20f42a5da4cde048419bac99db56498ca7/Untitled%2012.png)