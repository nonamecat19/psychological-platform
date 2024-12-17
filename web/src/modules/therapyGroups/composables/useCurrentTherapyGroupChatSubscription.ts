import { onMounted, onUnmounted } from 'vue'
import { socket } from '@core/lib/socket'
import { useCurrentTherapyGroupStore } from '../store/currentTherapyGroup.ts'

export function useCurrentTherapyGroupChatSubscription(chatId: number) {
  const event = `SEND_MESSAGE_GROUP`

  const store = useCurrentTherapyGroupStore()

  onMounted(() => {
    socket.connect()
    socket.on(`${event}:${chatId}`, (data: string) => {
      store.addMessage(JSON.parse(data))
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
