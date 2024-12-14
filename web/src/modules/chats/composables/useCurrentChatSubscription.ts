import { onMounted, onUnmounted } from 'vue'
import { socket } from '@core/lib/socket'
import { useCurrentChatStore } from '@chats/store/currentChat.ts'

export function useCurrentChatSubscription(chatId: number) {
  const event = `SEND_MESSAGE`

  const currentChatStore = useCurrentChatStore()

  onMounted(() => {
    socket.connect()
    console.log({ chatId })
    socket.on(`${event}:${chatId}`, (data: string) => {
      console.log({ data })
      currentChatStore.addMessage(JSON.parse(data))
    })
  })

  onUnmounted(() => {
    socket.off(event)
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
