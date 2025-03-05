// eslint-disable-next-line @typescript-eslint/no-var-requires
const { Builder, Browser, By, until } = require('selenium-webdriver');
// eslint-disable-next-line @typescript-eslint/no-var-requires
const { describe, it, before, after } = require('node:test');

var assert;

// eslint-disable-next-line @typescript-eslint/no-var-requires
(async () => {
    const chai = await import('chai');
    assert = chai.assert;
})();

describe('Login_EmailInput', async () => {
    let driver;

    before(async function () {
        driver = await new Builder().forBrowser(Browser.FIREFOX).build();
    });

    after(async function () {
        await driver.quit();
    });

    it('should_input_email', async function () {
        await driver.get('http://localhost:3000/log-in');

        await driver.findElement(By.name('email'));
    });

    it('should_input_password', async function () {
        await driver.get('http://localhost:3000/log-in');

        await driver.findElement(By.name('password'));
    });

    it('should_click_login_button', async function () {
        await driver.get('http://localhost:3000/log-in');

        await driver.findElement(By.css('button[type=submit]'));
    });

    it('login_with_valid_email', async function () {
        await driver.findElement(By.name('email')).sendKeys('user@example.com');
        await driver.findElement(By.name('password')).sendKeys('123456');
        await driver.findElement(By.css('button[type=submit]')).click();

        await driver.wait(until.elementLocated(By.id('error')), 10000);
        const text = await driver.findElement(By.id('error')).getText();
        assert.equal(text, '');
    });

    it('login_with_invalid_email', async function () {
        await driver.findElement(By.name('email')).sendKeys('invalid email');
        await driver.findElement(By.name('password')).sendKeys('123456');
        await driver.findElement(By.css('button[type=submit]')).click();

        await sleep(100);
        await driver.wait(until.elementLocated(By.id('error')), 10000);
        const text = await driver.findElement(By.id('error')).getText();
        assert.equal(text, 'Email không hợp lệ');
    });
});

/**
 *
 * @param {*} ms
 * @returns Promise
 */
function sleep(ms) {
    return new Promise((resolve) => {
        setTimeout(resolve, ms);
    });
}
