import { MPPage } from './app.po';

describe('mp App', () => {
  let page: MPPage;

  beforeEach(() => {
    page = new MPPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
