package anticaptcha

import (
	"errors"
)

type Error struct {
	ErrorID int `json:"errorId"`
	Error   error
}

func NewError() *Error {
	return &Error{}
}

func (e *Error) ErrMsg() error {
	if e.ErrorID == 0 {
		return e.Error
	}
	codes := make(map[int]string)
	codes[1] = "Account authorization key not found in the system"
	codes[2] = "No idle captcha workers are available at the moment, please try a bit later or try increasing your maximum bid here"
	codes[3] = "The size of the captcha you are uploading is less than 100 bytes."
	codes[4] = "The size of the captcha you are uploading is more than 500,000 bytes."
	codes[10] = "Account has zeo or negative balance"
	codes[11] = "Request with current account key is not allowed from your IP. Please refer to IP list section located here"
	codes[12] = "Captcha could not be solved by 5 different workers"
	codes[13] = "100% recognition feature did not work due to lack of amount of guess attempts"
	codes[14] = "Request to API made with method which does not exist"
	codes[15] = "Could not determine captcha file type by its exif header or image type is not supported. The only allowed formats are JPG, GIF, PNG"
	codes[16] = "Captcha you are requesting does not exist in your current captchas list or has been expired. Captchas are removed from API after 5 minutes after upload."
	codes[20] = "'comment' property is required for this request"
	codes[21] = "Your IP is blocked due to API inproper use. Check the reason at https://anti-captcha.com/panel/tools/ipsearch"
	codes[22] = "Task property is empty or not set in createTask method. Please refer to API v2 documentation."
	codes[23] = "Task type is not supported or inproperly printed. Please check \"type\" parameter in task object."
	codes[24] = "Some of the required values for successive user emulation are missing."
	codes[25] = "Could not connect to proxy related to the task, connection refused"
	codes[26] = "Could not connect to proxy related to the task, connection timeout"
	codes[27] = "Connection to proxy for task has timed out"
	codes[28] = "Proxy IP is banned by target service"
	codes[29] = "Task denied at proxy checking state. Proxy must be non-transparent to hide our server IP."
	codes[30] = "Recaptcha task timeout, probably due to slow proxy server or Google server"
	codes[31] = "Recaptcha server reported that site key is invalid"
	codes[32] = "Recaptcha server reported that domain for this site key is invalid"
	codes[33] = "Recaptcha server reported that browser user-agent is not compatible with their javascript"
	codes[34] = "Captcha provider server reported that additional variable token has been expired. Please try again with new token."
	codes[35] = "Proxy does not support transfer of image data from Google servers"
	codes[36] = "Proxy does not support long GET requests with length about 2000 bytes and does not support SSL connections"
	codes[37] = "Could not connect to Factory Server API within 5 seconds"
	codes[38] = "Incorrect Factory Server JSON response, something is broken"
	codes[39] = "Factory Server API did not send any errorId"
	codes[40] = "Factory Server API reported errorId != 0, check this error"
	codes[41] = "Some of the required property values are missing in Factory form specifications. Customer must send all required values."
	codes[42] = "Expected other type of property value in Factory form structure. Customer must send specified value type."
	codes[43] = "Factory control belong to another account, check your account key."
	codes[44] = "Factory Server general error code"
	codes[45] = "Factory Platform general error code."
	codes[46] = "Factory task lifetime protocol broken during task workflow."
	codes[47] = "Task not found or not available for this operation"
	codes[48] = "Factory is sandboxed, creating tasks is possible only by Factory owner. Switch it to production mode to make it available for other customers."
	codes[49] = "Proxy login and password are incorrect"
	codes[50] = "Customer did not enable Funcaptcha Proxyless tasks in Customers Area - API Settings. All customers must read terms, pass mini test and sign/accept the form before being able to use this feature."
	codes[51] = "Recaptcha was attempted to be solved as usual one, instead of invisible mode. Basically you don't need to do anything when this error occurs, just continue sending tasks with this domain. Our system will self-learn to solve recaptchas from this sitekey in invisible mode."
	codes[52] = "Could not load captcha provider widget in worker browser. Please try sending new task."
	e.Error = errors.New(codes[e.ErrorID])
	return e.Error
}

func (e *Error) ErrMsgRu() error {
	if e.ErrorID == 0 {
		return e.Error
	}
	codes := make(map[int]string)
	codes[1] = "Авторизационный ключ не существует в системе или имеет неверный формат (длина не равняется 32 байтам)"
	codes[2] = "Нет свободных работников в данный момент, попробуйте позже либо повысьте свою максимальную ставку здесь"
	codes[3] = "Размер капчи которую вы загружаете менее 100 байт"
	codes[4] = "Размер капчи которую вы загружаете более 500,000 байт"
	codes[10] = "Баланс учетной записи ниже нуля или равен нулю"
	codes[11] = "Запрос с этого IP адреса с текущим ключом отклонен. Управление доступом по IP находится здесь"
	codes[12] = "5 разных работников не смогли разгадать капчу, задание остановлено"
	codes[13] = "Не хватило заданного количества дублей капчи для функции 100% распознавания."
	codes[14] = "Запрос в API выполнен на несуществующий метод"
	codes[15] = "Формат капчи не распознан по EXIF заголовку либо не поддерживается. Допустимые форматы: JPG, GIF, PNG"
	codes[16] = "Капча с таким ID не была найдена в системе. Убедитесь что вы запрашиваете состояние капчи в течение 300 секунд после загрузки."
	codes[20] = "Отсутствует комментарий в параметрах рекапчи версии API 1"
	codes[21] = "Доступ к API с этого IP запрещен из-за большого количества ошибок. Узнать причину можно здесь."
	codes[22] = "Отсутствует задача в методе createTask."
	codes[23] = "Тип задачи не поддерживается или указан не верно."
	codes[24] = "Неполные или некорректные данные об эмулируемом пользователе. Все требуемые поля не должны быть пустыми."
	codes[25] = "Не удалось подключиться к прокси-серверу - отказ в подключении"
	codes[26] = "Таймаут подключения к прокси-серверу"
	codes[27] = "Таймаут операции чтения прокси-сервера."
	codes[28] = "Прокси забанен на целевом сервисе капчи"
	codes[29] = "Ошибка проверки прокси. Прокси должен быть не прозрачным, скрывать адрес конечного пользователя. В противном случае Google будет фильтровать запросы с IP нашего сервера. "
	codes[30] = "Таймаут загрузки скрипта рекапчи, проблема либо в медленном прокси, либо в медленном сервере Google"
	codes[31] = "Ошибка получаемая от сервера рекапчи. Неверный/невалидный sitekey."
	codes[32] = "Ошибка получаемая от сервера рекапчи. Домен не соответствует sitekey."
	codes[33] = "Для задачи используется User-Agent неподдерживаемого рекапчей браузера."
	codes[34] = "Провайдер капчи сообщил что дополнительный изменяющийся токен устарел. Попробуйте создать задачу еще раз с новым токеном."
	codes[35] = "Прокси не поддерживает передачу изображений с серверов Google"
	codes[36] = "Прокси не поддерживает длинные (длиной 2000 байт) GET запросы и не поддерживает SSL подключения"
	codes[37] = "Не смогли подключиться к API сервера фабрики в течени 5 секунд."
	codes[38] = "Неправильный JSON ответ фабрики, что-то сломалось."
	codes[39] = "API фабрики не вернул обязательное поле errorId"
	codes[40] = "Ожидали errorId = 0 в ответе API фабрики, получили другое значение."
	codes[41] = "Значения некоторых требуемых полей в запросе к фабрике отсутствуют. Клиент должен прислать все требуемы поля."
	codes[42] = "Тип значения не соответствует ожидаемому в структуре задачи фабрики. Клиент должен прислать значение с требуемым типом."
	codes[43] = "Доступ к управлению фабрикой принадлежит другой учетной записи. Проверьте свой ключ доступа."
	codes[44] = "Общий код ошибки сервера фабрики."
	codes[45] = "Общий код ошибки платформы."
	codes[46] = "Ошибка в протоколе во время выполнения задачи фабрики."
	codes[47] = "Задача не найдена или недоступна для этой операции."
	codes[48] = "Фабрика находится в режиме песочницы, создание задач доступно только для владельца фабрики. Переведите фабрику в боевой режим, чтобы сделать ее доступной для всех клиентов."
	codes[49] = "Заданы неверные логин и пароль для прокси"
	codes[50] = "Заказчик не включил тип задач Funcaptcha Proxyless в панели клиентов - Настройки API. Все заказчики должны прочитать условия, пройти мини тест и подписать/принять форму до того как смогут использовать данный тип задач."
	codes[51] = "Обнаружена попытка решить невидимую рекапчу в обычном режиме. В случае возникновения этой ошибки вам ничего не нужно предпринимать, наша система через некоторое время начнет решать задачи с этим ключом в невидимом режиме. Просто шлите еще задачи с тем же ключом и доменом."
	codes[52] = "Не удалось загрузить виджет капчи в браузере работника. Попробуйте прислать новую задачу."
	e.Error = errors.New(codes[e.ErrorID])
	return e.Error
}

func (e *Error) set(err error) {
	e.ErrorID = 0
	e.Error = err
}

func (e *Error) setString(err string) {
	e.ErrorID = 0
	e.Error = errors.New(err)
}
