## Watson Assistant

The IBM Watson™ Assistant service (formerly known as Watson Conversation) combines machine learning, natural language understanding, and integrated dialog tools to create conversation flows between your apps and your users. This application uses the [Watson Developer Cloud Go SDK](https://github.com/watson-developer-cloud/go-sdk).

With Watson Assistant you can create cognitive agents--virtual agents that combine machine learning, natural language understanding, and integrated dialog scripting tools to build outstanding projects, such as a chat room with an integrated Watson chatbot.
 
##  Credentials

###  LocalDevConfig

This is where your local configuration is stored for Watson Assistant.
```
{
  "watson_assistant_url": "{{url}}",
  "watson_assistant_username": "{{username}}",
  "watson_assistant_password": "{{password}}"
}
```

## Usage

Boilerplate code for initializing a client object for the Watson Assistant API is included inside `services/ service_watson_assistant.go`. The `services/services.go` file creates an instance of a `AssistantV1` client available for use anywhere in the application, which is preconfigured to use the credentials supplied. You will need to specify the `workspaceID` for the assistant workspace you have created, as each service integration can include multiple workspaces. The full documentation for the `IBM Watson Assistant` service for Go can [be found here](https://www.ibm.com/watson/developercloud/assistant/api/v1), but a small getting started example can be found below:

```go
  import (
    fmt
    "<application-name>/services"
    . "github.com/watson-developer-cloud/go-sdk/assistantV1"
  )

  func main () { 
    services.Init()

    assistant := services.WatsonAssistant
    workspaceID := "the-workspace-id-from-your-generated-service"

    input := InputData{Text:"hello"}
    messageOptions := NewMessageOptions(workspaceID).
      SetInput(input)

    msg, msgErr := assistant.Message(messageOptions)
    if msgErr != nil {
      fmt.Println(msgErr)
      return
    }

    msgResult := GetMessageResult(msg)
    if msgResult != nil {
      fmt.Println(msgResult)
    }
  }
```

## Extended Example

An extended example of basic usages can be found in the [Watson Developer Cloud GitHub Repo](https://github.com/watson-developer-cloud/go-sdk/tree/master/examples/assistantV1.go).

Reference the [ConversationV1 API to leverage its full power](https://www.ibm.com/watson/developercloud/assistant/api/v1).