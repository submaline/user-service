package user_service

import (
	"context"
	"fmt"
	"github.com/bufbuild/connect-go"
	userv1 "github.com/submaline/user-service/gen/user/v1"
)

type UserService struct{}

func (us *UserService) GetAccount(
	_ context.Context, request *connect.Request[userv1.GetAccountRequest]) (
	*connect.Response[userv1.GetAccountResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (us *UserService) UpdateAccount(
	_ context.Context, request *connect.Request[userv1.UpdateAccountRequest]) (
	*connect.Response[userv1.UpdateAccountResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (us *UserService) GetProfile(
	_ context.Context, request *connect.Request[userv1.GetProfileRequest]) (
	*connect.Response[userv1.GetProfileResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (us *UserService) UpdateProfile(
	_ context.Context, request *connect.Request[userv1.UpdateProfileRequest]) (
	*connect.Response[userv1.UpdateProfileResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (us *UserService) AddFriend(
	_ context.Context, request *connect.Request[userv1.AddFriendRequest]) (
	*connect.Response[userv1.AddFriendResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (us *UserService) GetFriends(
	_ context.Context, request *connect.Request[userv1.GetFriendsRequest]) (
	*connect.Response[userv1.GetFriendsResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}
func (us *UserService) BlockUser(
	_ context.Context, request *connect.Request[userv1.BlockUserRequest]) (
	*connect.Response[userv1.BlockUserResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (us *UserService) GetBlockingUsers(
	_ context.Context, request *connect.Request[userv1.GetBlockingUsersRequest]) (
	*connect.Response[userv1.GetBlockingUsersResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}

func (us *UserService) UnBlockUser(
	_ context.Context, request *connect.Request[userv1.UnBlockUserRequest]) (
	*connect.Response[userv1.UnBlockUserResponse], error) {
	return nil, connect.NewError(
		connect.CodeUnimplemented,
		fmt.Errorf(fmt.Sprintf("%s is unimplemented", request.Spec().Procedure)))
}
