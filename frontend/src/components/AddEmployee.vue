<template>
  <div class="app">
    <h1>Add Employee</h1>

    <form @submit.prevent="addEmployee" class="employee-form">
      <div class="form-group">
        <label for="name">Name</label>
        <input id="name" v-model="name" placeholder="Employee Name" required />
      </div>

      <div class="form-group">
        <label for="role">Role</label>
        <input id="role" v-model="role" placeholder="Role" required />
      </div>

      <div class="form-group">
        <label for="language">Language</label>
        <input
          id="language"
          v-model="language"
          placeholder="Programming Languages"
          required
        />
      </div>

      <div class="form-group">
        <label for="salary">Salary</label>
        <input
          id="salary"
          v-model.number="salary"
          type="number"
          placeholder="Salary"
          required
        />
      </div>

      <div class="button-group">
        <button type="submit" class="btn primary">Add Employee</button>
        <button type="button" class="btn secondary" @click="cancelForm">Cancel</button>
      </div>
    </form>

    <!-- Success / Error Message -->
    <div
      v-if="message"
      :class="{ success: success, error: !success }"
      class="message"
    >
      {{ message }}
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue"

const name = ref("")
const role = ref("")
const salary = ref("")
const language = ref("")

const message = ref("")
const success = ref(true)

async function addEmployee() {
  try {
    const res = await fetch("/api/employees", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({
        name: name.value,
        role: role.value,
        language: language.value,        
        salary: salary.value,
      }),
    })

    if (!res.ok) throw new Error("Failed to add employee")

    const emp = await res.json()

    // Reset form
    resetForm()

    message.value = `Employee ${emp.name} added successfully with ID ${emp.id}!`
    success.value = true
    setTimeout(() => (message.value = ""), 3000)
  } catch (err) {
    message.value = err.message
    success.value = false
    setTimeout(() => (message.value = ""), 3000)
  }
}

function cancelForm() {
  resetForm()
  message.value = "Form cleared."
  success.value = true
  setTimeout(() => (message.value = ""), 2000)
}

function resetForm() {
  name.value = ""
  role.value = ""
  language.value = ""  
  salary.value = ""
}
</script>

<style>
.app {
  max-width: 400px;
  margin: auto;
  padding: 20px;
  font-family: sans-serif;
  background: #e7e7e8ff;
  border-radius: 8px;
  box-shadow: 0px 2px 6px rgba(0, 0, 0, 0.1);
}

.employee-form {
  display: flex;
  flex-direction: column;
}

.form-group {
  display: flex;
  flex-direction: column;
  margin-bottom: 15px;
}

label {
  margin-bottom: 5px;
  font-weight: bold;
  color: #111010ff;
}

input {
  padding: 10px;
  font-size: 14px;
  border: 1px solid rgba(221, 221, 220, 1);
  border-radius: 6px;
}

.button-group {
  display: flex;
  justify-content: space-between;
  gap: 10px;
}

.btn {
  flex: 1;
  padding: 10px;
  font-size: 15px;
  cursor: pointer;
  border-radius: 6px;
  border: none;
  transition: 0.2s;
}

.btn.primary {
  background: #007bff;
  color: white;
}

.btn.primary:hover {
  background: #0056b3;
}

.btn.secondary {
  background: #f1f1f1;
  color: #333;
}

.btn.secondary:hover {
  background: #ddd;
}

.message {
  margin-top: 15px;
  padding: 12px;
  border-radius: 6px;
  font-weight: 500;
}
.success {
  background-color: #d4edda;
  color: #155724;
}
.error {
  background-color: #f8d7da;
  color: #721c24;
}
</style>
