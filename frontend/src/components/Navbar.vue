<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-light sticky-top">
    <!-- replace with container-fluid if we want left aligned links -->
    <div class="container">
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link
              class="nav-link"
              active-class="active"
              aria-current-value="page"
              :to="{ name: 'Home' }"
            >
              Files
            </router-link>
          </li>
          <li class="nav-item">
            <router-link
              class="nav-link"
              active-class="active"
              aria-current-value="page"
              :to="{ name: 'Devices' }"
            >
              Storage Devices
            </router-link>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>
