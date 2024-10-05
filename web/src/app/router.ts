import {createRouter, createWebHistory} from 'vue-router'
import {supportRouter} from "../modules/support";
import {chatRouter} from "../modules/chats";

const routes = [
  chatRouter,
  supportRouter
]

export const appRouter = createRouter({
  history: createWebHistory(),
  routes,
})
