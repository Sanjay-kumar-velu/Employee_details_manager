import { createRouter, createWebHistory } from "vue-router"
import AddEmployee from "./components/AddEmployee.vue"
import EmployeeList from "./components/EmployeeList.vue"

const routes = [
  { path: "/", redirect: "/add" },
  { path: "/add", component: AddEmployee },
  { path: "/list", component: EmployeeList }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
