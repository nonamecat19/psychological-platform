import {createRouter, createWebHistory} from 'vue-router'
import {supportRouter} from "@support";
import {chatRouter} from "@chats";

const routes = [
  chatRouter,
  supportRouter
]

export const appRouter = createRouter({
  history: createWebHistory(),
  routes,
})
