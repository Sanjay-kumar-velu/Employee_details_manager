<template>
  <div id="app">
    <!-- Navbar -->
    <nav class="navbar">
      <h1 class="logo">Employee Details Manager</h1>
      <div class="links">
        <router-link to="/add" class="nav-link">Add Employee</router-link>
        <router-link to="/list" class="nav-link">Employee List</router-link>
      </div>
    </nav>

    <!-- Page Content -->
    <main class="container">
      <router-view />
    </main>
  </div>
</template>

<style>
/* Reset some defaults */
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

/* Navbar styling */
.navbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: #007bff;
  padding: 15px 30px;
  color: white;
}

.logo {
  font-size: 20px;
  font-weight: bold;
}

.links {
  display: flex;
  gap: 15px;
}

.nav-link {
  color: white;
  text-decoration: none;
  font-weight: 500;
  padding: 8px 12px;
  border-radius: 6px;
  transition: background 0.3s ease;
}

.nav-link:hover {
  background: rgba(255, 255, 255, 0.2);
}

.router-link-exact-active {
  background: white;
  color: #007bff;
  font-weight: bold;
}

/* Main container for page content */
.container {
  max-width: 900px;
  margin: 20px auto;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 10px;
  box-shadow: 0px 2px 8px rgba(0, 0, 0, 0.1);
}
</style>
