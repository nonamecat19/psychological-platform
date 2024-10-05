import ChatView from "./views/ChatView.vue"

export const chatRouter = {
  path: "/chat",
  children: [
    {
      path: "",
      component: ChatView,
    },
  ],
}
