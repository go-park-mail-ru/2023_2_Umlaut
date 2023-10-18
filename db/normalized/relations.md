## Диаграммы

[ER диаграмма в draw.io](Umlaut.drawio)

[Mermaid описание](https://mermaid.live/edit#pako:eNqVlNFuozAQRX_F8nPTD8gbLd4UlQDCbqVKkZATZhOrgCPbqK0S_n3tQEMgaTfhBeveM55hPHiHVzIHPMWgfMHXipeLCtnnhZIU7feTidyhx3iehF4QMTRFAc1SksQpI-m1oN-CR3PC3hJyKeTJo5nzRjzatYJ7RGVQ8pwFfi_pkheF01PYSmWo4abWvW1ECdrwcovY96o1m5P6f89g4NOgiJdwIemLBjWDKgd1wUwU_D23D9v5oFdKbI2Q1cgJpXwX1Xqkkrxe8Qv0k1wuBeiRGpR8DQk3m17PuQH0IJTZ5Pyrl5dSFiiubLkwaMvotG5q0PcWKfFJxAIvRM3-_n6_a3vdnnNvnsFX5JpzUYykhGv9IVU-kikvzPl5d8Pn23TxzA1r9BqHr4Rmf4KUsswx19CUPMaRf4J3yCBgMieUejNiAy3NbE9p1kmDoJ8_uznljttd0ybQ2g4Cs-sbf4juC8Lg2dXtXn728PZDYwYUi0-og_OfQo_l2JGbB_a6yDw2KIh5M9TIQ6ZugNwrszLtgdsGFN_hElTJRW5vvkPoApsNWB9P7TLn6n2BF5XjeG0k_apWeGpUDXe43ro_qbsrW7H5B0cZgh8)
## Описание таблиц

1. **Таблица "COMPLAINT" (Жалобы):**
    -  Эта таблица хранит информацию о жалобах, поданных пользователями. В ней фиксируются жалобы с указанием текста сообщения, времени подачи и связанных сущностей, таких как пользователь-жалующийся, пользователь-нарушитель и тип жалобы.

2. **Таблица "COMPLAINT-TYPE" (Типы жалоб):**
    -  Эта таблица содержит перечень различных типов жалоб, которые могут быть использованы в таблице "COMPLAINT" для классификации жалоб.

3. **Таблица "USER" (Пользователи):**
    -  Эта таблица содержит информацию о пользователях приложения. Здесь сохраняются данные о пользовательских профилях, включая имя, пол, возраст, интересы, теги, путь к изображению профиля, дату рождения и другие характеристики.

4. **Таблица "CREDENTIAL" (Учетные данные):**
    -  В этой таблице хранятся учетные данные пользователей, включая адрес электронной почты (Mail), хэш пароля (Password) и соль (Salt) для обеспечения безопасности авторизации.

5. **Таблица "DIALOG" (Диалоги):**
    -  Эта таблица предназначена для хранения информации о диалогах между пользователями. Здесь фиксируются идентификаторы диалогов, а также связи с пользователями и последним сообщением в диалоге.

6. **Таблица "MESSAGE" (Сообщения в диалогах):**
    -  В этой таблице хранятся сообщения, отправленные в диалогах между пользователями. Для каждого сообщения указывается текст сообщения, дата отправки и связь с соответствующим диалогом и отправителем.

7. **Таблица "LIKE" (Лайки):**
    -  Эта таблица используется для отслеживания действий "лайк" в приложении. Она включает в себя информацию о лайках, включая дату и связи с пользователями, которые поставили лайк и пользователями, которым поставили лайк.

8. **Таблица "Tag" (Теги):**
   -  Эта таблица хранит список тегов, которые могут быть привязаны к пользователям.

9. **Таблица "UserTag" (Теги пользователей):**
   -  Эта таблица представляет связь между пользователями и тегами. Она указывает, какие теги связаны с конкретными пользователями.

## Функциональные зависимости

**Relation "COMPLAINT" (Жалобы):**

{PK ID} -> {Message, ReportedID, ComplaintTypeID, Timestamp}

**Relation "USER" (Пользователи):**

{PK ID} -> {Name, UserGender, Looking, Hobbies, ImagePath, Birthday}

**Relation "COMPLAINT-TYPE" (Типы жалоб):**

{PK ID} -> {Name}

**Relation "CREDENTIAL" (Учетные данные):**

{PK ID} -> {Mail, Password, Salt}

**Relation "DIALOG" (Диалоги):**

{PK ID} -> {FK_User1ID, FK_User2ID}

**Relation "MESSAGE" (Сообщения в диалогах):**

{PK ID} -> {Message Text, Date, FK_DialogID, FK_SenderID}

**Relation "LIKE" (Лайки):**

{PK ID} -> {Timestamp, FK_LikedByUserID, FK_LikedToUserID}

**Relation "Tag" (Теги):**

{PK ID} -> {Name}

**Relation "UserTag" (Теги пользователей):**

{PK ID} -> {UsersID, TagID}