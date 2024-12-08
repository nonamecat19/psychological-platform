import { createRouter, createWebHistory } from 'vue-router'
import { chatRouter } from '@chats'
import { supportRouter } from '@support'
import { adminRouter } from '../modules/admin'
import { authRouter } from '../modules/auth'
import { usersRouter } from '../modules/users'

const routes = [chatRouter, supportRouter, adminRouter, authRouter, usersRouter]

export const appRouter = createRouter({
  history: createWebHistory(),
  routes,
})
