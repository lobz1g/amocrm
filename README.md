# AmoCRM

[![Go Report Card](https://goreportcard.com/badge/github.com/lobz1g/amocrm)](https://goreportcard.com/report/github.com/lobz1g/amocrm)

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

#### Получить компанию/сделку/задачу по ID
```go
   api := amocrm.NewAmo("YOUR_LOGIN", "YOUR_API_KEY", "YOUR_DOMAIN")

   comp, err := api.Company.Id(123456)
   if err != nil {
      log.Println(err)
   }
   fmt.Println(comp)
```

#### Добавление компании/сделки/задачи
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

#### Обновление компании/сделки/задачи
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

### Компания
#### Получить все компании по отвественному
```go
    api := amocrm.NewAmo("YOUR_LOGIN", "YOUR_API_KEY", "YOUR_DOMAIN")
    
    companies, err := api.Company.Responsible(11234)
    if err != nil {
      log.Println(err)
    }
    
    for _, value := range companies {
      fmt.Println(value)
    }
```

### Сделка
#### Получить все сделки по отвественному
```go
    api := amocrm.NewAmo("YOUR_LOGIN", "YOUR_API_KEY", "YOUR_DOMAIN")
    
    leads, err := api.Lead.Responsible(11234)
    if err != nil {
      log.Println(err)
    }
    
    for _, value := range leads {
      fmt.Println(value)
    }
```
#### Получить все сделки по статусу
```go
    api := amocrm.NewAmo("YOUR_LOGIN", "YOUR_API_KEY", "YOUR_DOMAIN")
    
    leads, err := api.Lead.Status(1123456)
    if err != nil {
      log.Println(err)
    }
    
    for _, value := range leads {
      fmt.Println(value)
    }
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

