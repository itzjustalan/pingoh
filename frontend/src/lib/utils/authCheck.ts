import { redirect } from "@tanstack/react-router"
import { authStore } from "../stores/auth"

export const checkAuth = () => {
    console.log("eeeeeeeee")
    console.log("eeeeeeeee", authStore.getState().user)
    if (authStore.getState().user === undefined) {
      throw redirect({
        to: "/auth/signin",
        search: {
          redirect: location.href,
        },
      })
    }
}
