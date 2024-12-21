<script setup lang="ts">
const props = defineProps({
  text: {
    type: String,
  },
  isMine: {
    type: Boolean,
    default: false,
  },
  userName: {
    type: String,
  },
})
</script>

<template>
  <div
    :class="[
      'p-2 rounded-md',
      props.isMine ? 'bg-white text-black' : 'bg-zinc-600 text-white',
    ]"
  >
    <div class="text-emerald-500">
      {{ userName }}
    </div>
    <div>
      {{ text }}
    </div>
  </div>
</template>
