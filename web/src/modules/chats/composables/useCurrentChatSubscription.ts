import { onMounted, onUnmounted } from 'vue'
import { socket } from '@core/lib/socket'
import { useCurrentChatStore } from '@chats/store/currentChat.ts'

export function useCurrentChatSubscription(chatId: number) {
  const event = `SEND_MESSAGE_PRIVATE`

  const currentChatStore = useCurrentChatStore()

  onMounted(() => {
    socket.connect()
    socket.on(`${event}:${chatId}`, (data: string) => {
      currentChatStore.addMessage(JSON.parse(data))
    })
  })

  onUnmounted(() => {
    socket.off(`${event}:${chatId}`)
    socket.disconnect()
  })

  function sendMessage(message: string) {
    socket.emit(
      event,
      JSON.stringify({
        token: localStorage.getItem('token'),
        message,
        chatId,
      }),
    )
  }

  return {
    sendMessage,
  }
}
