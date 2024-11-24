import { createRouter, createWebHistory } from 'vue-router'
import { chatRouter } from '@chats'
import { supportRouter } from '@support'
import { adminRouter } from '../modules/admin'
import { authRouter } from '../modules/auth'

const routes = [chatRouter, supportRouter, adminRouter, authRouter]

export const appRouter = createRouter({
  history: createWebHistory(),
  routes,
})
