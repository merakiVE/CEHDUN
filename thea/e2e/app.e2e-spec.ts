import { FendPage } from './app.po';

describe('thea App', () => {
  let page: FendPage;

  beforeEach(() => {
    page = new FendPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
