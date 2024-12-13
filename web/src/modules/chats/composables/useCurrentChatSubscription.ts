import { onMounted, onUnmounted } from 'vue'
import { socket } from '@core/lib/socket'

export function useCurrentChatSubscription(chatId: number) {
  const event = `CUSTOM_EVENT`

  // const auctionsStore = useAuctionsStore()

  onMounted(() => {
    socket.connect()
    socket.on(event, (data) => {
      // auctionsStore.addNewAuctionBid(auctionId, data)
      console.log({ data })
    })
  })

  onUnmounted(() => {
    socket.off(event)
    socket.disconnect()
  })

  function sendMessage(message: string) {
    socket.emit(event, message)
  }

  return {
    sendMessage,
  }
}
