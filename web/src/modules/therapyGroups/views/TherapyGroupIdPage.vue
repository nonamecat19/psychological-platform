<script setup lang="ts">
import { useRoute } from 'vue-router'
import { computed, onMounted, ref } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import ChatMessage from '@chats/components/ChatMessage.vue'
import { useCurrentUserStore } from '@core/stores/useCurrentUserStore.ts'
import { useCurrentTherapyGroupChatSubscription } from '../composables/useCurrentTherapyGroupChatSubscription.ts'
import { useCurrentTherapyGroupStore } from '../store/currentTherapyGroup.ts'

const route = useRoute()

const store = useCurrentTherapyGroupStore()

const id = computed(() => route.params.id)

onMounted(() => {
  store.setCurrentTherapyGroupId(Number(id.value))
  store.initCurrentTherapyGroupId()
})

const { sendMessage } = useCurrentTherapyGroupChatSubscription(+id.value)

async function handleSendMessage() {
  sendMessage(inputValue.value)
  inputValue.value = ''
}

const currentUserStore = useCurrentUserStore()

const inputValue = ref<string>('')
</script>
<template>
  <div class="h-full w-full flex flex-col">
    <div class="grow p-4 flex flex-col gap-2">
      <ChatMessage
        v-for="message of store?.data?.Messages"
        :text="message.Content"
        :is-mine="message.UserID === currentUserStore.data.ID"
      />
    </div>
    <div class="p-2 pt-0 w-full flex gap-2">
      <InputText v-model="inputValue" class="w-full" />
      <Button @click="handleSendMessage">
        <i class="pi pi-send" />
      </Button>
    </div>
  </div>
</template>
