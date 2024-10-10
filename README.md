﻿# Go Cars Orders
Change env based on your system logic
 
    SQL_HOST=localhost
    SQL_USER=root
    SQL_PASSWORD=root1234
    SQL_DB=golang-test
    SQL_PORT=3306

﻿# API Go Cars Orders

[GET] http://localhost:3400/api/car

Response:

    {
        "status_code": 200,
        "message": "Get data success",
        "data": [
            {
                "car_id": 2,
                "car_name": "SDDSSDDSSD",
                "day_rate": 123123312132,
                "month_rate": 288381283,
                "image": "SDDSSDDSSD.com"
            }
        ]
    }

[POST] http://localhost:3400/api/car

Request:

    {
        "car_name": "rizkiTEST123",
        "day_rate": 99999,
        "month_rate": 123132132,
        "image": "rizkiTest123.com"
    }

Response:

    {
        "status_code": 200,
        "message": "Insert data cars success"
    }

[DELETE] http://localhost:3400/api/car/:id 

Response:
    
    {
        "status_code": 200,
        "message": "Delete data cars by 3 is success!"
    }

[PUT] http://localhost:3400/api/car/:id 

Request: 

    {
        "car_name": "BudiTEST01239",
        "day_rate": 123666,
        "month_rate": 123666,
        "image": "BudiTEST01239.com"
    }

Response:
    
    {
        "status_code": 200,
        "message": "Update data cars by 3 success!"
    }
