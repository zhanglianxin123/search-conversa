namespace go api.gpt
struct ConversaReq {
    1: string Id; // 添加 api 注解为方便进行参数绑定
    2: string Content;
}

struct ConversaResp {
    1: string Id;
    2: Message Message;
}
struct Message {
   1: string Role;
   2: string Content;
}

service ConversaService {
    ConversaResp CreateConversaMethod() (api.post="/api/gpt/conversation");
    ConversaResp ConversaMethod(1: ConversaReq request) (api.put="/api/gpt/conversation");
}