<template>
  <div class="app">
    <h1>Employee List</h1>
    <table>
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Role</th>
          <th>Language</th>
          <th>Salary</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="emp in employees" :key="emp.id">
          <td>{{ emp.id }}</td>
          <td>{{ emp.name }}</td>
          <td>{{ emp.role }}</td>
          <td>{{ emp.language }}</td>
          <td>{{ emp.salary }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue"

const employees = ref([])

async function fetchEmployees() {
  const res = await fetch("/api/employees")
  employees.value = await res.json()
}

onMounted(fetchEmployees)
</script>

<style>

.app {

  width: 80%;
  max-width: 500px;     /* increase from content-fit to larger */
  margin: 0 auto;        /* keep it centered */
}

table {
  width: 100%;  
  border-collapse: collapse;
  margin-top: 25px;
  font-family: Arial, sans-serif;
  box-shadow: 0px 2px 8px rgba(5, 27, 18, 0.1);
  border-radius: 6px;
  overflow: hidden;
}

th, td {
  padding: 12px;
  text-align: left;
}

th {
  background-color: #007bff; /* blue header */
  color: white;
  font-weight: bold;
}

tr:nth-child(even) {
  background-color: #f2f8ff; /* light blue */
}

tr:nth-child(odd) {
  background-color: #ffffff; /* white */
}

tr:hover {
  background-color: #dceeff; /* slightly darker blue on hover */
}
</style>
