import { createRouter, createWebHistory } from "vue-router"
import { chatRouter } from "@chats"
import { supportRouter } from "@support"

const routes = [chatRouter, supportRouter]

export const appRouter = createRouter({
  history: createWebHistory(),
  routes,
})
