// eslint-disable-next-line @typescript-eslint/no-var-requires
const { Builder, Browser, By, Key, until } = require('selenium-webdriver');
// eslint-disable-next-line @typescript-eslint/no-var-requires
const { describe, it, before, after } = require('node:test');

var assert;
var expect;

// eslint-disable-next-line @typescript-eslint/no-var-requires
(async () => {
    const chai = await import('chai');
    assert = chai.assert;
    expect = chai.expect;
})();

describe('Login_Success', async () => {
    let driver;

    before(async function () {
        driver = await new Builder().forBrowser(Browser.FIREFOX).build();
    });

    after(async function () {
        await driver.quit();
    });

    it('should login', async function () {
        await driver.get('http://localhost:3000/log-in');

        await driver
            .findElement(By.name('email'))
            .sendKeys('mostima@mail.com', Key.RETURN);
        await driver
            .findElement(By.name('password'))
            .sendKeys('12345678', Key.RETURN);

        await driver.wait(
            until.urlContains('http://localhost:3000/dashboard/root'),
            10000
        );
    });
});

describe('Login_FailInvalidPass', async () => {
    let driver;

    before(async function () {
        driver = await new Builder().forBrowser(Browser.FIREFOX).build();
    });

    after(async function () {
        await driver.quit();
    });

    it('should fail to login with invalid password', async function () {
        await driver.get('http://localhost:3000/log-in');

        await driver.findElement(By.name('email')).sendKeys('mostima@mail.com');
        await driver
            .findElement(By.name('password'))
            .sendKeys('wrongpass', Key.RETURN);

        // Đợi thông báo lỗi hiển thị
        await driver.wait(
            until.elementLocated(By.id('error')), // Giả sử có class này cho thông báo lỗi
            10000
        );

        // Kiểm tra vẫn ở trang đăng nhập
        const currentUrl = await driver.getCurrentUrl();
        assert.strictEqual(
            currentUrl.includes('http://localhost:3000/log-in'),
            true
        );
    });
});

describe('Login_FailInvalidEmail', async () => {
    let driver;

    before(async function () {
        driver = await new Builder().forBrowser(Browser.FIREFOX).build();
    });

    after(async function () {
        await driver.quit();
    });

    it('should fail to login with notfound email', async function () {
        await driver.get('http://localhost:3000/log-in');

        await driver
            .findElement(By.name('email'))
            .sendKeys('notfound@mail.com');
        await driver
            .findElement(By.name('password'))
            .sendKeys('password123', Key.RETURN);

        // Đợi thông báo lỗi hiển thị
        await driver.wait(
            until.elementLocated(By.id('error')), // Giả sử có class này cho thông báo lỗi
            10000
        );

        // Kiểm tra vẫn ở trang đăng nhập
        const currentUrl = await driver.getCurrentUrl();
        assert.strictEqual(
            currentUrl.includes('http://localhost:3000/log-in'),
            true
        );
    });
});

describe('Login_SuccessWithTimeout', async () => {
    let driver;

    before(async function () {
        driver = await new Builder().forBrowser(Browser.FIREFOX).build();

        await driver.get('http://localhost:3000/dashboard/root');

        await driver.wait(
            until.urlContains('http://localhost:3000/log-in'),
            10000
        );
    });

    after(async function () {
        await driver.quit();
    });

    it('should login with timeout', async function () {
        await driver.get('http://localhost:3000/log-in');

        await driver
            .findElement(By.name('email'))
            .sendKeys('mostima@mail.com', Key.RETURN);
        await driver
            .findElement(By.name('password'))
            .sendKeys('12345678', Key.RETURN);

        // Đợi chuyển trang 2 giây
        await driver.wait(
            until.urlContains('http://localhost:3000/dashboard/root'),
            2000
        );
    });
});

describe('Login_FailWithFiveTimes', async () => {
    let driver;

    before(async function () {
        driver = await new Builder().forBrowser(Browser.FIREFOX).build();

        await driver.get('http://localhost:3000/dashboard/root');

        await driver.wait(
            until.urlContains('http://localhost:3000/log-in'),
            10000
        );
    });

    after(async function () {
        await driver.quit();
    });

    it('should fail login 5 times and show error', async function () {
        await driver.get('http://localhost:3000/log-in');

        for (let i = 0; i < 5; i++) {
            // Nhập email
            await driver.findElement(By.name('email')).clear();
            await driver
                .findElement(By.name('email'))
                .sendKeys('mostima@mail.com');

            // Nhập mật khẩu sai
            await driver.findElement(By.name('password')).clear();
            await driver
                .findElement(By.name('password'))
                .sendKeys('wrongPass', Key.RETURN);

            // Đợi thông báo lỗi xuất hiện
            const errMgs = driver.findElement(By.id('error'));
            await driver.wait(
                until.elementTextContains(errMgs, 'Invalid credentials.'),
                5000
            );
        }

        // Nhập email
        await driver.findElement(By.name('email')).clear();
        await driver.findElement(By.name('email')).sendKeys('mostima@mail.com');

        // Nhập mật khẩu sai
        await driver.findElement(By.name('password')).clear();
        await driver
            .findElement(By.name('password'))
            .sendKeys('wrongPass', Key.RETURN);

        // Đợi thông báo lỗi xuất hiện
        const errMgs = driver.findElement(By.id('error'));
        await driver.wait(
            until.elementTextContains(
                errMgs,
                'Tài khoản đã bị khóa do nhập sai quá nhiều lần'
            ),
            5000
        );
    });
});
