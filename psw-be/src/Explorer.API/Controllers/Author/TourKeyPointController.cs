﻿using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Tours.API.Dtos;
using Explorer.Tours.API.Public;
using Explorer.Tours.API.Public.Administration;
using Explorer.Tours.Core.UseCases;
using Explorer.Tours.Core.UseCases.Administration;
using Microsoft.AspNetCore.Authorization;
using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.RazorPages;
using System.Text;
using System.Text.Json;
namespace Explorer.API.Controllers.Author
{
    //[Authorize(Policy = "authorPolicy")]
    [Route("api/tourKeyPoint")]
    public class TourKeyPointController : BaseApiController
    {
        private readonly ITourKeyPointService _tourKeyPointService;
        private readonly IPublicTourKeyPointService _publicTourKeyPointService;
        public TourKeyPointController(ITourKeyPointService tourKeyPointService, IPublicTourKeyPointService publicTourKeyPointService)
        {
            _tourKeyPointService = tourKeyPointService;
            _publicTourKeyPointService = publicTourKeyPointService;
        }

        [HttpGet]
        public ActionResult<PagedResult<TourKeyPointDto>> GetAll([FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _tourKeyPointService.GetPaged(page, pageSize);
            return CreateResponse(result);
        }

        /*[HttpGet("tour/{tourId:int}")]
        public ActionResult<PagedResult<TourKeyPointDto>> GetByTourId(int tourId)
        {
            var result = _tourKeyPointService.GetByTourId(tourId);
            return CreateResponse(result);
        }*/

        [HttpGet("tour/{tourId:int}")]
        public async Task<ActionResult<List<TourKeyPointDto>>> GetByTourIdAsync(int tourId)
        {
            try
            {
                var result = await _tourKeyPointService.GetByTourIdAsync(tourId);
                return CreateResponse(result);
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }

        [HttpGet("{id:int}")]
        public ActionResult<TourKeyPointDto> Get(int id)
        {
            var result = _tourKeyPointService.Get(id);
            return CreateResponse(result);
        }

        /*[HttpPost]
        public Task<string> Create([FromBody] TourKeyPointDto tourKeyPoint)
        {
            var result = _tourKeyPointService.CreateAsync(tourKeyPoint, _httpClient);
            return result;
        }*/

        [HttpPost]
        public async Task<ActionResult<TourKeyPointDto>> CreateAsync([FromBody] TourKeyPointDto tourKeyPoint)
        {
            try
            {
                var createdTourKeypoint = await _tourKeyPointService.CreateAsync(tourKeyPoint);
                return CreateResponse(createdTourKeypoint);
            }
            catch (Exception e)
            {
                return StatusCode(500, e.Message);
            }
        }

        [HttpPut("{id:int}")]
        public ActionResult<TourKeyPointDto> Update([FromBody] TourKeyPointDto tourKeyPoint)
        {
            var result = _tourKeyPointService.Update(tourKeyPoint);
            return CreateResponse(result);
        }

        [HttpDelete("{id:int}")]
        public ActionResult Delete(int id)
        {
            var result = _tourKeyPointService.Delete(id);
            return CreateResponse(result);

        }

        //public async Task<string> CreateAsync(TourKeyPointDto tourKeypointDto, HttpClient _httpClient)
        //{
        //    using StringContent jsonContent = new(JsonSerializer.Serialize(tourKeypointDto), Encoding.UTF8, "application/json");
        //    using HttpResponseMessage response = await _httpClient.PostAsync("http://localhost:8080/keypoints", jsonContent);
        //    response.EnsureSuccessStatusCode();
        //    var jsonResponse = await response.Content.ReadAsStringAsync();
        //    return jsonResponse;
        //}

        [HttpGet("public")]
        public ActionResult<PagedResult<PublicTourKeyPointDto>> GetAllPublic([FromQuery] int page, [FromQuery] int pageSize)
        {
            var result = _publicTourKeyPointService.GetPaged(page, pageSize);
            return CreateResponse(result);
        }

        [HttpPost("public")]
        public ActionResult<PublicTourKeyPointDto> CreatePublic([FromBody] PublicTourKeyPointDto tourKeyPoint)
        {
            var result = _publicTourKeyPointService.Create(tourKeyPoint);
            return CreateResponse(result);
        }

        [HttpPut("public/{tourId}/{status}")]
        public ActionResult<PublicTourKeyPointDto> ChangeStatus(int tourId, String status)
        {
            var result = _publicTourKeyPointService.ChangeStatus(tourId, status);
            return CreateResponse(result);
        }

        [HttpGet("public/{status}")]
        public ActionResult<PagedResult<PublicTourKeyPointDto>> GetByStatus(String status)
        {
            var result = _publicTourKeyPointService.GetByStatus(status);
            return CreateResponse(result);
        }

    }
}
