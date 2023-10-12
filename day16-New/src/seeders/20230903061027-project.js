'use strict';

/** @type {import('sequelize-cli').Migration} */
module.exports = {
  async up (queryInterface, Sequelize) {
    await queryInterface.bulkInsert('projects', [
      {
        name: 'Project-1 tes',
        start_date: '2023/09/03',
        end_date: '2023/09/24',
        description:'Tes project-1 apakah bisa atau tidak',
        technologies: [true,false,true,true],
        image:'coba.jpg',
        userid:1,
        createdAt: new Date(),
        updatedAt: new Date()
      },
      {
        name: 'Project-2 tes',
        start_date: '2023/09/03',
        end_date: '2023/09/24',
        description:'Tes project-2 apakah bisa juga atau tidak',
        technologies: [false,false,true,true],
        image:'coba.jpg',
        userid:2,
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
