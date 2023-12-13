let socket = new WebSocket("wss://umlaut-bmstu.me/websocket");
//let socket = new WebSocket("ws://localhost:8000/api/v1/ws/messenger");

socket.onopen = () => {
    console.log("Successfully Connected");
};

socket.onmessage = function (event) {
    console.log(event.data);
}

socket.onclose = event => {
    console.log("Socket Closed Connection: ", event);
};

socket.onerror = error => {
    console.log("Socket Error: ", error);
};


// структура сообщения:
// type Message struct {
// 	Id          int       `json:"id"`
// 	SenderId    int       `json:"sender_id"`
// 	RecipientId int       `json:"recipient_id"`
// 	DialogId    int       `json:"dialog_id"`
// 	Text        string    `json:"message_text"`
// 	CreatedAt   time.Time `json:"created_at"`
// }

// type Message struct {
//     type string `json:"type" // "Message" || "Match"`
//     payload interface `json:"payload"` // объект сообщения или объект пары
// }

// примеры создания сообщения:
socket.send(`{
    "sender_id": 101,
    "recipient_id": 201,
    "dialog_id": 123,
    "message_text": "Привет, как дела?"
}`)
// примеры редактирования:
socket.send(`{
    "id": 1,
    "sender_id": 101,
    "recipient_id": 201,
    "dialog_id": 123,
    "message_text": "Привет, как дела?",
    "is_read": true
}`)
