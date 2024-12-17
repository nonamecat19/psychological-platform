<script setup lang="ts">
const props = defineProps({
  text: {
    type: String,
  },
  isMine: {
    type: Boolean,
    default: false,
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
    {{ text }}
  </div>
</template>
