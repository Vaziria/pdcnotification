export interface INotif {
    action: 'add_token' | 'send_action'
    email: string
    message: string
    tokens: string[]
}