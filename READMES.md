# Dokumentasi API Task Manager

## 1. Get All Tasks
**URI**: `/tasks`  
**Metode**: `GET`  
**Request Body**: Tidak ada  
**Response Body**:  
**Status Code**: `200 OK`  
```json
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

## 2. Get Task by ID
**URI**: `/task/{id}`  
**URI PARAMETER**: `{ID}` wajib  
**Metode**: `GET`  
**Request Body**: Tidak ada  
**Response Body**:  
**Status Code**: `200 OK`  
```json
{
    "id": 1,
    "datetime": "08-10-2024 14:00",
    "description": "Example task 1",
    "status": "Created"
}
**Status Code**: `404 Not Found`  
```json
{
    "error": "Task not found"
} // jika ID yang diberikan tidak sesuai dengan yang ada

## 3. Create Task
**URI**: `/task`  
**Metode**: `POST`  
**Request Body**:  
```json
{
    "datetime": "2024-10-09T14:00:00Z",
    "description": "New task description",
    "status": "Created" // Allowed values: "Created", "Pending", "Completed"
}

**atau**
```json
{
    "description": "New task description",
    "status": "Created" // Allowed values: "Created", "Pending", "Completed"
}

**Response Body**:  
**Status Code**: `201 Created`  
```json
{
    "id": 4,
    "datetime": "09-10-2024 14:00", // Waktu yang diinput jika diisi ataupun waktu saat mengirim request jika dikosongkan
    "description": "New task description",
    "status": "Created"
}
**Status Code**: `400 Bad Request` 
```json
{
    "error": "Invalid status. Allowed values are 'Created', 'Pending', or 'Completed'."
} // jika status yang diinput tidak sesuai

## 4. Update Task Status
**URI**: `/task/{id}`  
**URI PARAMETER**: `{ID}` wajib  
**Metode**: `PUT`  
**Request Body**: Tidak ada  
**Response Body**:  
**Status Code**: `200 OK`  
```json
{
    "id": 1,
    "datetime": "08-10-2024 14:00",
    "description": "Example task 1",
    "status": "Pending"
} // jika Task sbebelumnya berstatus Created maka akan berubah menjadi Pending
**atau**
```json
{
    "id": 1,
    "datetime": "08-10-2024 14:00",
    "description": "Example task 1",
    "status": "Completed"
} // jika Task sbebelumnya berstatus Pending maka akan berubah menjadi Completed
**Status Code**: `404 Not Found`  
```json
{
    "error": "Task not found"
} // jika ID yang diberikan tidak sesuai dengan yang ada
**Status Code**: `403 Forbidden`  
```json
{
    "error": "Task is already completed and cannot be updated"
} // jika task yang di update sudah berstatus completed

## 5. Delete Task
**URI**: `/task/{id}`  
**URI PARAMETER**: `{ID}` wajib  
**Metode**: `DELETE`  
**Request Body**: Tidak ada  
**Response Body**:  
**Status Code**: `200 OK`  
```json
{
    "message": "Task deleted successfully"
}
**Status Code**: `404 Not Found`  
```json
{
    "error": "Task not found"
} // jika ID yang diberikan tidak sesuai dengan yang ada