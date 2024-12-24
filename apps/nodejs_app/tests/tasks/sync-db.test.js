const { syncDB } = require("../../tasks/sync-db");

describe('pruebas en SyncDB', () => {
    test('deberia haber marcado 2', () => {
        syncDB()
        const times = syncDB()

        console.log('se llamo', times);

        expect(times).toBe(2)
    })
})