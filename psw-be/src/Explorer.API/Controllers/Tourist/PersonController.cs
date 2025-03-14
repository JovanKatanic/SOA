﻿using Explorer.Blog.API.Dtos;
using Explorer.Stakeholders.API.Dtos;
using Explorer.Stakeholders.API.Public;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Internal;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;

namespace Explorer.API.Controllers.Tourist
{
    [Authorize(Policy = "touristPolicy")]
    [Route("api/tourist/person")]
    public class PersonController : BaseApiController
    {
        private readonly IPersonService _personService;
        public PersonController(IPersonService personService)
        {
            _personService = personService;
        }

        [HttpGet("{id:int}")]
        public ActionResult<PersonDto> Get(int id)
        {
            var result = _personService.Get(id);
            return CreateResponse(result);
        }

        [HttpGet]
        public ActionResult<PersonDto> GetAuthorsAndTourists()
        {
            var result = _personService.GetAuthorsAndTourists();
            return CreateResponse(result);
        }

        [AllowAnonymous]
        [HttpGet("followers/{id:int}")]
        public ActionResult<List<PersonDto>> GetAllFollowers(int id) 
        {
            var result = _personService.GetAllFollowers(id);
            return CreateResponse(result);
        }



        //[AllowAnonymous]
        //[HttpGet("followings/{id:int}")]
        //public ActionResult<List<PersonDto>> GetAllFollowings(int id)
        //{
        //    var result = _personService.GetAllFollowings(id);
        //    return CreateResponse(result);
        //}

        [AllowAnonymous]
        [HttpGet("followings/{id:int}")]
        public async Task<ActionResult<List<PersonDto>>> GetAllFollowingsAsync(int id)
        {
            try
            {
                var followings = await _personService.GetAllFollowingsAsync(id);
                return CreateResponse(followings);
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }

        [AllowAnonymous]
        [HttpGet("recommendedfollowings/{id:int}")]
        public async Task<ActionResult<List<PersonDto>>> GetRecommendedFollowingsAsync(int id)
        {
            try
            {
                var followings = await _personService.GetRecommendedFollowingsAsync(id);
                return CreateResponse(followings);
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }

        [HttpPut("{id:int}")]
        public ActionResult<PersonDto> Update([FromBody] PersonDto person)
        {
            var result = _personService.Update(person);
            return CreateResponse(result);
        }
    }
}
