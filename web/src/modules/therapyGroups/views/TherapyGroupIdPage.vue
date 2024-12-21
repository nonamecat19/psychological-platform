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

const currentGroupStore = useCurrentTherapyGroupStore()
const currentUserStore = useCurrentUserStore()

const id = computed(() => route.params.id)

onMounted(() => {
  currentGroupStore.setCurrentTherapyGroupId(Number(id.value))
  currentGroupStore.initCurrentTherapyGroupId()
})

const { sendMessage } = useCurrentTherapyGroupChatSubscription(+id.value)

async function handleSendMessage() {
  sendMessage(inputValue.value)
  inputValue.value = ''
}

const inputValue = ref<string>('')
</script>
<template>
  <div class="h-full w-full flex flex-col">
    <div class="grow p-4 flex flex-col gap-2">
      <!--      <pre>-->
      <!--      {{ currentGroupStore.data?.Messages }}}-->
      <!--      </pre>-->

      <ChatMessage
        v-for="message of currentGroupStore?.data?.Messages"
        :text="message.Content"
        :is-mine="message.UserID === currentUserStore?.data?.ID"
        :user-name="
          message.User
            ? `${message.User.Name} ${message.User.Surname} ${message.User.Role === 'psychologist' ? '(Specialist)' : ''}`
            : 'Anonymous'
        "
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
