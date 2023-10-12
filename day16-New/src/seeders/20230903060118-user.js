'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface, Sequelize) {
    
    await queryInterface.bulkInsert('users', [
      {
        name: 'Adrian Saputra',
        email: 'adriansaputra@gmail.com',
        password:'1234',
        createdAt: new Date(),
        updatedAt: new Date()
      },
      {
        name: 'Ivan Swarna',
        email: 'ivanswarna@gmail.com',
        password:'1234',
        createdAt: new Date(),
        updatedAt: new Date()
      }
    ], {});
    
  },

  async down (queryInterface, Sequelize) {
    /**
     * Add commands to revert seed here.
     *
     * Example:
     * await queryInterface.bulkDelete('People', null, {});
     */
  }
};
