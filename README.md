Dokumentasi API Task Manager

1. Get All Tasks
URI: /tasks
Metode: GET
Request Body: Tidak ada
Response Body:
    Status Code: 200 OK
    Body:
        [
            {
                "id": 1,
                "datetime": "08-10-2024 14:00",
                "description": "Example task 1",
                "status": "Created"
            },
            {
                "id": 2,
                "datetime": "09-10-2024 12:30",
                "description": "Example task 2",
                "status": "Pending"
            },
            {
                "id": 3,
                "datetime": "09-10-2024 16:30",
                "description": "Example task 3",
                "status": "Completed"
            }
        ]

 
2. Get Task by ID
URI: /task/{id}
URI PARAMETER: {ID} wajib
Metode: GET
Request Body: Tidak ada
Response Body:
    Status Code: 200 OK
    Body:
        {
        "id": 1,
        "datetime": "08-10-2024 14:00",
        "description": "Example task 1",
        "status": "Created"
        }
    Status Code: 404 Not Found (Jika ID di bawah 0 dan di atas jumlah task)
    Body:
        {
        "error": "Task not found"
        }

3. Create Task
URI: /task
Metode: POST
Request Body:
    {
    "datetime": "2024-10-09T14:00:00Z",
    "description": "New task description",
    "status": "Created" // Allowed values: "Created", "Pending", "Completed"
    }

    atau

    {
    "description": "New task description",
    "status": "Created" // Allowed values: "Created", "Pending", "Completed"
    } // datetime akan di isi waktu saat mengirim request


Response Body:
    Status Code: 201 Created
    Body:
        {
        "id": 4,
        "datetime": "09-10-2024 14:00", // Waktu yang diinput jika diisi ataupun waktu saat mengirim request jika dikosongkan
        "description": "New task description",
        "status": "Created"
        }
        
    Status Code: 400 Bad Request (Jika Status yang diinput tidak sesuai)
    Body:
        {
        "error": "Invalid status. Allowed values are 'Created', 'Pending', or 'Completed'."
        }

4. Update Task Status
URI: /task/{id}
URI PARAMETER: {ID} wajib
Metode: PUT
Request Body: Tidak ada
Response Body:
    Status Code: 200 OK
    Body:
        {
        "id": 1,
        "datetime": "08-10-2024 14:00",
        "description": "Example task 1",
        "status": "Pending"
        } // Jika sebelumnya Created akan berubah menjadi Pending

        atau

        {
        "id": 1,
        "datetime": "08-10-2024 14:00",
        "description": "Example task 1",
        "status": "Completed"
        } // Jika sebelumnya Pending akan berubah menjadi Completed

    Status Code: 404 Not Found (Jika ID di bawah 0 dan di atas jumlah task)
    Body:
        {
        "error": "Task not found"
        }
        
    Status Code: 403 Forbidden (Jika Task yang diupdate sudah berstatus Completed)
    Body:
        {
        "error": "Task is already completed and cannot be updated"
        }

5. Delete Task
URI: /task/{id}
URI PARAMETER: {ID} wajib
Metode: DELETE
Request Body: Tidak ada
Response Body:
    Status Code: 200 OK
    Body:
        {
        "message": "Task deleted successfully"
        }

    Status Code: 404 Not Found
    Body:
        {
        "error": "Task not found"
        }