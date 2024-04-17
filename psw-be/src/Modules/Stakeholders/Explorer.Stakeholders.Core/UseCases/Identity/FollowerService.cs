using AutoMapper;
using Explorer.BuildingBlocks.Core.Domain;
using Explorer.BuildingBlocks.Core.UseCases;
using Explorer.Stakeholders.API.Dtos;
using Explorer.Stakeholders.API.Public;
using Explorer.Stakeholders.API.Public.Identity;
using Explorer.Stakeholders.Core.Domain;
using Explorer.Stakeholders.Core.Domain.Followers;
using Explorer.Stakeholders.Core.Domain.RepositoryInterfaces;
using FluentResults;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Text;
using System.Text.Json;
using System.Threading.Tasks;

namespace Explorer.Stakeholders.Core.UseCases.Identity
{
    public class FollowerService : BaseService<FollowerDto, Follower>, IFollowerService
    {
        private readonly IFollowerRepository _followerRepository;
        private readonly IPersonRepository _personRepository;
        private readonly HttpClient _httpClient;

        public FollowerService(IFollowerRepository followerRepository, IMapper mapper, 
            IPersonRepository personRepository, HttpClient _httpClient) : base(mapper)
        {
            _followerRepository = followerRepository;
            _personRepository = personRepository;
            this._httpClient = _httpClient;
        }

        private static HttpClient sharedClient = new()
        {
            BaseAddress = new Uri("http://follower_service:8082")
        };
        public Result<List<SavedNotificationDto>> GetFollowersNotifications(int id)
        {
            try
            {
                var followers = _followerRepository.GetFollowersNotifications(id);
                List<Person> people = _personRepository.GetAllFollowers(id);

                List<SavedNotificationDto> saveNotifications = new List<SavedNotificationDto>();

                foreach(var follower in followers)
                {
                    var person = people.Find(p => p.Id == follower.FollowerId);

                    saveNotifications.Add(new SavedNotificationDto { FullName = person.Name + " " + person.Surname, TimeOfArrival = follower.Notification.TimeOfArrival });
                }

                return saveNotifications;
            }
            catch (KeyNotFoundException e)
            {
                return Result.Fail(FailureCode.NotFound).WithError(e.Message);
            }
        }
        public Result<FollowerDto> Create(FollowerDto follower)
        {
            try
            {
                var result = _followerRepository.Create(MapToDomain(follower));
                return MapToDto(result);
            }
            catch (ArgumentException e)
            {
                return Result.Fail(FailureCode.InvalidArgument).WithError(e.Message);
            }
        }

        public async Task<FollowerDto> CreateAsync(FollowerDto follower)
        {
            using StringContent jsonContent = new(System.Text.Json.JsonSerializer.Serialize(follower), Encoding.UTF8, "application/json");

            using HttpResponseMessage response = await sharedClient.PostAsync("/createFollowe", jsonContent);

            response.EnsureSuccessStatusCode();

            var jsonString = await response.Content.ReadAsStringAsync();
            var jsonObject = JsonDocument.Parse(jsonString).RootElement;

            FollowerDto followerDto = new FollowerDto
            {
                Id = jsonObject.GetProperty("id").GetInt32(),
                FollowerId = jsonObject.GetProperty("followerId").GetInt32(),
                FollowedId = jsonObject.GetProperty("followedId").GetInt32(),
                Notification = ParseNotification(jsonObject.GetProperty("notification"))
            };

            return followerDto;
        }

        private FollowerNotificationDto ParseNotification(JsonElement notificationElement)
        {
            if (notificationElement.ValueKind == JsonValueKind.Object)
            {
                return JsonSerializer.Deserialize<FollowerNotificationDto>(notificationElement.GetRawText());
            }

            return new FollowerNotificationDto();
        }

        public Result Delete(int followerId, int followedId)
        {
            try
            {
                _followerRepository.Delete(followerId, followedId);
                return Result.Ok();
            }
            catch (KeyNotFoundException e)
            {
                return Result.Fail(FailureCode.NotFound).WithError(e.Message);
            }
        }

        public Result<List<FollowerDto>> GetFollowings(int id)
        {
            var result = _followerRepository.GetFollowings(id);
            return MapToDto(result);
        }
    }
}
