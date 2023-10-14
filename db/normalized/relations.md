## Диаграммы

[ER диаграмма в draw.io](https://drive.google.com/file/d/13nEG0242lOdOozOpXfhOZTflDkC_dgan/view?usp=sharing)

[Mermaid описание](https://mermaid.live/edit#pako:eNqVlNFugjAUhl-l6bW8AHcojRIRCK0mJiTkKB02Cpi2ZjPqu6-o09HNhfWG04-_p-dvD5zwuik4djGXvoBSQpXVyIxRPEtCL4gYOp8dpzmhOSUpclFA85QkccpI2lfovxA-gMOWCTFLJh7N29DWn26gHaLWKJnmgf9Emn9oNONKQcmftADNERMVVxqq_Y1fbo9rgT1yRlBxC80Vl2NeF1x2l3ulrQybZivq0qKTZrUSXFmUQWmjoDJuEtAby9FQSL0p4NgxZB3kv6x9pUiJTyIWeKF1je2dPF_-EPe5GhA7CyWg1HsjCwtT2OlOVb7ZIh7bjRUt4nBBaN7Of9PdZs6MUOqN274axREzx0PzO-oseu2gU8QjXf9mZCbu1ZBhMCWWyRb5-XD5zeRV9ffuL_bAA1xxWYEozHd-zZBhveGmBbBrwgLkNsNZfTE6OOiGHus1drU88AE-7Nuc9z8Ddt9gp_jlE94WL_U)

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

## Функциональные зависимости

**Relation "COMPLAINT" (Жалобы):**

{PK ID, ReporterID} -> {Message, ReportedID, ComplaintTypeID, Timestamp}

**Relation "USER" (Пользователи):**

{PK ID} -> {Name, UserGender, Age, Looking, Hobbies, Tags, ImagePath, Birthday}

**Relation "COMPLAINT-TYPE" (Типы жалоб):**

{PK ID} -> {Name}

**Relation "CREDENTIAL" (Учетные данные):**

{PK ID} -> {Mail, Password, Salt}

**Relation "DIALOG" (Диалоги):**

{PK ID} -> {FK_User1ID, FK_User2ID, FK_LastMessageID}

**Relation "MESSAGE" (Сообщения в диалогах):**

{PK ID} -> {Message Text, Date, FK_DialogID, FK_SenderID}

**Relation "LIKE" (Лайки):**

{PK ID} -> {Timestamp, FK_LikedByUserID, FK_LikedToUserID}
