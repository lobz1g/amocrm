# AmoCRM

Клиент для работы с api AmoCRM

## Установка

```go
go get -u github.com/lobz1g/amocrm
```

## Примеры работы

### Компании/сделки/задачи

#### Базовая работа с компанией/сделкой/задачей
```go
   api := amocrm.NewAmo("YOUR_LOGIN", "YOUR_API_KEY", "YOUR_DOMAIN")

   allCompanies, err := api.Company.All()
   if err != nil {
      log.Println(err)
   }

   for _, value := range allCompanies {
      fmt.Println(value)
   }
```

#### Получить компанию по ID
```go
   api := amocrm.NewAmo("YOUR_LOGIN", "YOUR_API_KEY", "YOUR_DOMAIN")

   comp, err := api.Company.Id(123456)
   if err != nil {
      log.Println(err)
   }
   fmt.Println(comp)
```

#### Добавление компании
```go
    api := amocrm.NewAmo("YOUR_LOGIN", "YOUR_API_KEY", "YOUR_DOMAIN")
    
    comp := api.Company.Create()
    comp.Name = "test"
    id, err := api.Company.Add(comp)
    if err != nil {
        log.Println(err)
    }
    fmt.Println(id)
```

#### Обновление компании
```go
    api := amocrm.NewAmo("YOUR_LOGIN", "YOUR_API_KEY", "YOUR_DOMAIN")
    
    comp, err := api.Company.Id(123456)
    if err != nil {
        log.Println(err)
    }
    comp.Name="another test"
    err=api.Company.Update(comp)
    if err != nil {
        log.Println(err)
    }
```

### Аккаунт
#### Получить всю информацию об аккаунте
```go
    api := amocrm.NewAmo(login, key, domain)
    
    acc, err := api.Account.Get()
    if err != nil {
        log.Println(err)
    }
    fmt.Println(acc)
```
### Задачи
#### Закрыть задачу
```go
    api := amocrm.NewAmo(login, key, domain)
    
    task, err := api.Task.Id(123456)
    if err != nil {
        log.Println(err)
    }
    task.Result.Text = "close task"
    err = api.Task.Close(task)
    if err != nil {
        log.Println(err)
    }
```
### Примечание
#### Создать примечание
```go
    note := api.Note.Create()
    note.NoteType = 1
    note.ElementId = 123456
    note.ElementType = 2
    note.Text = "test note"
    id, err := api.Note.Add(note)
    if err != nil {
        log.Println(err)
    }
    fmt.Println(id)
```

