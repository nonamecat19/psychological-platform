import { fileURLToPath, URL } from 'url'
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
// @ts-expect-error
import eslint from 'vite-plugin-eslint'
import vueDevTools from 'vite-plugin-vue-devtools'

export default defineConfig({
  plugins: [vue(), vueDevTools()],
  resolve: {
    alias: [
      {
        find: '@app',
        replacement: fileURLToPath(new URL('./src/app', import.meta.url)),
      },
      {
        find: '@core',
        replacement: fileURLToPath(new URL('./src/core', import.meta.url)),
      },
      {
        find: '@chats',
        replacement: fileURLToPath(
          new URL('./src/modules/chats', import.meta.url),
        ),
      },
      {
        find: '@therapyGroups',
        replacement: fileURLToPath(
          new URL('./src/modules/therapyGroups', import.meta.url),
        ),
      },
      {
        find: '@support',
        replacement: fileURLToPath(
          new URL('./src/modules/support', import.meta.url),
        ),
      },
    ],
  },
})
