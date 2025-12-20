function isAuthenticated() {
    // TODO: reemplazar por lógica real (Pinia + token, etc.)
    return !!localStorage.getItem('access_token')
}

router.beforeEach((to, from, next) => {
    const loggedIn = isAuthenticated()

    if (to.meta.requiresAuth && !loggedIn) {
        return next({ name: 'login', query: { redirect: to.fullPath } })
    }

    if (to.meta.guestOnly && loggedIn) {
        return next({ name: 'dashboard' })
    }

    return next()
})

export default router
